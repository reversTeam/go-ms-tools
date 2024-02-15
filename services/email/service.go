package email

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	"github.com/reversTeam/go-ms-tools/pkg/scylla"
	"github.com/reversTeam/go-ms-tools/services/abs"
	ms "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	pb "github.com/reversTeam/go-ms-tools/services/email/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Define the service structure
type Service struct {
	*abs.Service
	pb.UnimplementedEmailServer
	scyllaGlobal *scylla.ScyllaDBManager
}

func NewService(ctx *core.Context, name string, config core.ServiceConfig) *Service {
	var dbConf scylla.DatabaseConfig
	err := mapstructure.Decode(config.Config["databases"], &dbConf)
	if err != nil {
		panic(err)
	}

	scyllaGlobalSession, err := scylla.NewScyllaDBManager("global", dbConf.Scylla["global"])
	if err != nil {
		panic(err)
	}

	return &Service{
		Service:      abs.NewService(ctx, name, config),
		scyllaGlobal: scyllaGlobalSession,
	}
}

func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterEmailHandlerFromEndpoint(gh.Ctx.Main, gh.Mux, endpoint, gh.Grpc.Opts)
}

func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterEmailServer(gs.Server, o)
}

func (o *Service) GetClient(conn *grpc.ClientConn) any {
	return pb.NewEmailClient(conn)
}

func (o *Service) List(in *empty.Empty, stream pb.Email_ListServer) error {
	_, scyllaSpan := core.Trace(stream.Context(), "email", "List")
	defer scyllaSpan.End()

	q := o.scyllaGlobal.Get("SELECT people_id, email, status FROM emails")

	r := pb.EmailResponse{}
	for q.Scan(&r.PeopleId, &r.Email, &r.Status) { // Adapter les champs
		if err := stream.Send(&r); err != nil {
			return err
		}
	}

	if err := q.Close(); err != nil {
		return err
	}

	return nil
}

func (o *Service) Get(ctx context.Context, in *pb.EmailEntity) (*pb.EmailResponse, error) {
	_, span := core.Trace(ctx, "email", "Get")
	defer span.End()

	var emailResponse pb.EmailResponse
	req := o.scyllaGlobal.GetOne("SELECT people_id, email, status FROM emails WHERE people_id = ? AND email = ?", in.PeopleId, in.Email)
	err := req.Scan(&emailResponse.PeopleId, &emailResponse.Email, &emailResponse.Status)

	return &emailResponse, err
}

func (o *Service) Create(ctx context.Context, in *pb.EmailCreateParams) (*pb.EmailResponse, error) {
	if err := o.scyllaGlobal.ExecuteQuery("INSERT INTO email (people_id, email, status) VALUES (?, ?, ?)", in.PeopleId, in.Email, in.Status); err != nil {
		return nil, err
	}

	return &pb.EmailResponse{
		PeopleId: in.PeopleId,
		Email:    in.Email,
		Status:   in.Status,
	}, nil
}

func (o *Service) Update(ctx context.Context, in *pb.EmailUpdateParams) (*ms.Response, error) {
	if err := o.scyllaGlobal.ExecuteQuery("UPDATE email SET status = ?, validated_at = ?, expired_at = ? WHERE people_id = ? AND email = ?",
		in.Status, in.ValidatedAt, in.ExpiredAt, in.PeopleId, in.Email); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Email updated successfully",
	}, nil
}

func (o *Service) Delete(ctx context.Context, in *pb.EmailEntity) (*ms.Response, error) {
	if err := o.scyllaGlobal.ExecuteQuery("DELETE FROM email WHERE people_id = ? AND email = ?", in.PeopleId, in.Email); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Email deleted successfully",
	}, nil
}
