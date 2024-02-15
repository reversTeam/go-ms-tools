package signin

import (
	"context"
	"fmt"
	"time"

	"log"

	"github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
	"github.com/reversTeam/go-ms-tools/pkg/scylla"
	"github.com/reversTeam/go-ms-tools/services/abs"
	pbAbs "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	pbAccount "github.com/reversTeam/go-ms-tools/services/account/protobuf"
	pbEmail "github.com/reversTeam/go-ms-tools/services/email/protobuf"
	pbPeople "github.com/reversTeam/go-ms-tools/services/people/protobuf"
	pb "github.com/reversTeam/go-ms-tools/services/signin/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
)

// Service structure for signin
type Service struct {
	*abs.Service
	pb.UnimplementedSigninServer
	scyllaGlobal *scylla.ScyllaDBManager
	scyllaAuth   *scylla.ScyllaDBManager
}

// NewService creates a new signin service
func NewService(ctx *core.Context, name string, config core.ServiceConfig) *Service {
	var dbConf scylla.DatabaseConfig
	err := mapstructure.Decode(config.Config["databases"], &dbConf)
	if err != nil {
		panic(err)
	}

	scyllaGlobal, err := scylla.NewScyllaDBManager("auth", dbConf.Scylla["global"])
	if err != nil {
		panic(err)
	}
	scyllaAuth, err := scylla.NewScyllaDBManager("auth", dbConf.Scylla["auth"])
	if err != nil {
		panic(err)
	}

	return &Service{
		Service:      abs.NewService(ctx, name, config),
		scyllaGlobal: scyllaGlobal,
		scyllaAuth:   scyllaAuth,
	}
}

func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterSigninHandlerFromEndpoint(gh.Ctx.Main, gh.Mux, endpoint, gh.Grpc.Opts)
}

func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterSigninServer(gs.Server, o)
}

func (o *Service) GetClient(conn *grpc.ClientConn) any {
	return pb.NewSigninClient(conn)
}

func (o *Service) Register(ctx context.Context, req *pb.RegisterRequest) (*pbAbs.Response, error) {
	id, err := gocql.RandomUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	expiredAt := now.Add(15 * 24 * time.Hour) // Ajoute 15 jours à la date/heure actuelle

	if req.Email == "" || req.Password == "" {
		return nil, fmt.Errorf("missing required arguments")
	}

	// Génération d'un validation_token aléatoire de 128 caractères
	validationToken, err := core.GenerateRandomString(128)
	if err != nil {
		return nil, fmt.Errorf("failed to generate validation token: %v", err)
	}

	// Vérification s'il existe déjà un signin non expiré pour cet email
	var existingSigninCount int
	if err := o.scyllaAuth.GetOne("SELECT COUNT(id) FROM signin WHERE email = ? AND expired_at > ? ALLOW FILTERING", req.Email, now).Scan(&existingSigninCount); err != nil {
		log.Printf("error checking for existing signin : %s", err)
	}

	if existingSigninCount > 0 {
		return nil, fmt.Errorf("a signin request for this email already exists and is not expired")
	}

	// Cryptage du mot de passe
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt password: %v", err)
	}

	// Insertion de la demande de signin avec expired_at et validation_token
	if err := o.scyllaAuth.ExecuteQuery("INSERT INTO signin (id, created_at, updated_at, email, firstname, lastname, birthday, password, validation_token, status, expired_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		id, now, now, req.Email, req.Firstname, req.Lastname, req.Birthday, hashedPassword, validationToken, "pending", expiredAt); err != nil {
		return nil, fmt.Errorf("failed to insert signin request: %v", err)
	}

	return &pbAbs.Response{
		Message: "Your subscription is registered, please check your email to confirm your account",
	}, nil
}

func (o *Service) Validate(ctx context.Context, req *pb.ValidateSigninRequest) (*pbAbs.Response, error) {
	var signin struct {
		ID              gocql.UUID
		Status          string
		ExpiredAt       time.Time
		ValidationToken string
		Firstname       string
		Lastname        string
		Birthday        string
		Password        string
		Email           string
	}

	if err := o.scyllaAuth.GetOne(`SELECT id, status, firstname, lastname, birthday, email, password, expired_at, validation_token FROM signin WHERE id = ?`, req.Id).Scan(
		&signin.ID,
		&signin.Status,
		&signin.Firstname,
		&signin.Lastname,
		&signin.Birthday,
		&signin.Email,
		&signin.Password,
		&signin.ExpiredAt,
		&signin.ValidationToken,
	); err != nil {
		return nil, fmt.Errorf("error querying signin by ID: %v", err)
	}

	// Vérification que le signin n'est pas déjà validé
	if signin.Status == "validated" {
		return nil, fmt.Errorf("signin is already validated")
	}

	// Vérification que l'expiration_at n'est pas dépassé
	if time.Now().After(signin.ExpiredAt) {
		return nil, fmt.Errorf("signin is expired")
	}

	// Vérification que le validation_token est correct
	if signin.ValidationToken != req.ValidationToken {
		return nil, fmt.Errorf("invalid validation token")
	}

	iPeopleGrpcClient, err := o.ClientManager.GetClient("people")
	if err != nil {
		return nil, err
	}
	peopleGrpcClient := iPeopleGrpcClient.(pbPeople.PeopleClient)
	peopleResponse, err := peopleGrpcClient.(pbPeople.PeopleClient).Create(ctx, &pbPeople.PeopleCreateParams{
		Firstname: signin.Firstname,
		Lastname:  signin.Lastname,
		Birthday:  signin.Birthday,
	})
	if err != nil {
		log.Printf("BIRTHDAY PEOPLE: %s\n", signin.Birthday)
		log.Printf("ERROR WHEN CREATE PEOPLE: %s\n", err)
		return nil, err
	}

	iEmailGrpcClient, err := o.ClientManager.GetClient("email")
	if err != nil {
		return nil, err
	}
	emailGrpcClient := iEmailGrpcClient.(pbEmail.EmailClient)
	emailResponse, err := emailGrpcClient.Create(ctx, &pbEmail.EmailCreateParams{
		PeopleId: peopleResponse.Id,
		Email:    signin.Email,
		Status:   "validated",
	})
	if err != nil {
		log.Printf("ERROR WHEN CREATE EMAIL: %s\n", err)
		return nil, err
	}

	iAccountGrpcClient, err := o.ClientManager.GetClient("account")
	if err != nil {
		return nil, err
	}
	accountGrpcClient := iAccountGrpcClient.(pbAccount.AccountClient)
	accountResponse, err := accountGrpcClient.Create(ctx, &pbAccount.AccountCreateParams{
		PeopleId: emailResponse.PeopleId,
		Password: signin.Password,
		Status:   "validated",
	})
	if err != nil {
		log.Printf("ERROR WHEN CREATE ACCOUNT: %s\n", err)
		return nil, err
	}

	// Mise à jour du statut de validation en base
	if err := o.scyllaAuth.ExecuteQuery(`UPDATE signin SET status = 'validated', validated_at = ?, account_id = ? WHERE id = ?`, time.Now(), accountResponse.PeopleId, req.Id); err != nil {
		return nil, fmt.Errorf("error updating signin validation status: %v", err)
	}

	return &pbAbs.Response{
		Message: "Your account has been confirm",
	}, nil
}
