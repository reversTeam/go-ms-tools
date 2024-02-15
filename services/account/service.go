package account

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	"github.com/reversTeam/go-ms-tools/pkg/scylla"
	"github.com/reversTeam/go-ms-tools/services/abs"
	ms "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	pb "github.com/reversTeam/go-ms-tools/services/account/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Define the service structure
type Service struct {
	*abs.Service
	pb.UnimplementedAccountServer
	scyllaAuth *scylla.ScyllaDBManager
}

func NewService(ctx *core.Context, name string, config core.ServiceConfig) *Service {
	var dbConf scylla.DatabaseConfig
	err := mapstructure.Decode(config.Config["databases"], &dbConf)
	if err != nil {
		panic(err)
	}

	scyllaAuthSession, err := scylla.NewScyllaDBManager("auth", dbConf.Scylla["auth"])
	if err != nil {
		panic(err)
	}

	return &Service{
		Service:    abs.NewService(ctx, name, config),
		scyllaAuth: scyllaAuthSession,
	}
}

func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterAccountHandlerFromEndpoint(gh.Ctx.Main, gh.Mux, endpoint, gh.Grpc.Opts)
}

func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterAccountServer(gs.Server, o)
}

func (o *Service) GetClient(conn *grpc.ClientConn) any {
	return pb.NewAccountClient(conn)
}

func (o *Service) List(in *empty.Empty, stream pb.Account_ListServer) error {
	_, scyllaSpan := core.Trace(stream.Context(), "account", "List")
	defer scyllaSpan.End()

	q := o.scyllaAuth.Get("SELECT people_id, status, password, validated_at, expired_at, signin_id FROM account")

	r := pb.AccountResponse{}
	for q.Scan(&r.PeopleId, &r.Status, &r.Password, &r.ValidatedAt, &r.ExpiredAt, &r.SigninId) {
		if err := stream.Send(&r); err != nil {
			return err
		}
	}

	if err := q.Close(); err != nil {
		return err
	}

	return nil
}

func (o *Service) Get(ctx context.Context, in *pb.AccountEntity) (*pb.AccountResponse, error) {
	_, span := core.Trace(ctx, "account", "Get")
	defer span.End()

	var account pb.AccountResponse
	req := o.scyllaAuth.GetOne("SELECT people_id, status, password, validated_at, expired_at, signin_id FROM account WHERE people_id = ?", in.PeopleId)
	err := req.Scan(&account.PeopleId, &account.Status, &account.Password, &account.ValidatedAt, &account.ExpiredAt, &account.SigninId)

	return &account, err
}

func (o *Service) Create(ctx context.Context, in *pb.AccountCreateParams) (*pb.AccountResponse, error) {
	peopleId, err := gocql.RandomUUID()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if err := o.scyllaAuth.ExecuteQuery("INSERT INTO account (people_id, signin_id, created_at, updated_at, validated_at, email, password, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", peopleId, in.SigninId, now, now, now, in.Email, in.Password, in.Status); err != nil {
		log.Printf("ACCOUNT CREATION ON ERROR IN THE MS : %s", err)
		return nil, err
	}

	return &pb.AccountResponse{
		PeopleId: peopleId.String(),
		Password: in.Password,
		Status:   in.Status,
	}, nil
}

func (o *Service) Update(ctx context.Context, in *pb.AccountUpdateParams) (*ms.Response, error) {
	if err := o.scyllaAuth.ExecuteQuery("UPDATE account SET password = ?, status = ?, email = ?, validated_at = ?, updated_at = ?, expired_at = ?, signin_id = ? WHERE people_id = ?",
		in.Password, in.Status, in.Email, in.ValidatedAt, time.Now(), in.ExpiredAt, in.SigninId, in.PeopleId); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Account updated successfully",
	}, nil
}

func (o *Service) Delete(ctx context.Context, in *pb.AccountEntity) (*ms.Response, error) {
	if err := o.scyllaAuth.ExecuteQuery("DELETE FROM account WHERE people_id = ?", in.PeopleId); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Account deleted successfully",
	}, nil
}
