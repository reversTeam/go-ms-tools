package people

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/mitchellh/mapstructure"
	"github.com/reversTeam/go-ms-tools/pkg/scylla"
	"github.com/reversTeam/go-ms-tools/services/abs"
	ms "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	pb "github.com/reversTeam/go-ms-tools/services/people/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Service struct {
	*abs.Service
	pb.UnimplementedPeopleServer
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
	return pb.RegisterPeopleHandlerFromEndpoint(gh.Ctx.Main, gh.Mux, endpoint, gh.Grpc.Opts)
}

func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterPeopleServer(gs.Server, o)
}

func (o *Service) GetClient(conn *grpc.ClientConn) any {
	return pb.NewPeopleClient(conn)
}

func (o *Service) List(in *empty.Empty, stream pb.People_ListServer) error {
	_, scyllaSpan := core.Trace(stream.Context(), "people", "List")
	defer scyllaSpan.End()

	q := o.scyllaGlobal.Get("SELECT id, status, firstname, lastname, birthday FROM people")

	r := pb.PeopleResponse{}
	id := &gocql.UUID{}
	for q.Scan(id, &r.Status, &r.Firstname, &r.Lastname, &r.Birthday) { // Adapter les champs
		r.Id = id.String()
		if err := stream.Send(&r); err != nil {
			return err
		}
	}

	if err := q.Close(); err != nil {
		return err
	}

	return nil
}

func (o *Service) Get(ctx context.Context, in *ms.EntityRequest) (*pb.PeopleResponse, error) {
	id, err := gocql.ParseUUID(in.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	_, span := core.Trace(ctx, "people", "Get")
	defer span.End()

	var people pb.PeopleResponse
	req := o.scyllaGlobal.GetOne("SELECT id, firstname, lastname, birthday FROM people WHERE id = ?", id)
	err = req.Scan(&people.Id, &people.Firstname, &people.Lastname, &people.Birthday)

	return &people, err
}

func (o *Service) Create(ctx context.Context, in *pb.PeopleCreateParams) (*pb.PeopleResponse, error) {
	id, err := gocql.RandomUUID()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	if err := o.scyllaGlobal.ExecuteQuery("INSERT INTO people (id, created_at, updated_at, status, firstname, lastname, birthday) VALUES (?, ?, ?, ?, ?, ?, ?)", id, now, now, "validated", in.Firstname, in.Lastname, in.Birthday); err != nil {
		return nil, err
	}

	return &pb.PeopleResponse{
		Id:        id.String(),
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Birthday:  in.Birthday,
	}, nil
}

func (o *Service) Update(ctx context.Context, in *pb.PeopleUpdateParams) (*ms.Response, error) {
	if err := o.scyllaGlobal.ExecuteQuery("UPDATE people SET updated_at = ?, firstname = ?, lastname = ?,  birthday = ? WHERE id = ?",
		time.Now(), in.Firstname, in.Lastname, in.Birthday, in.Id); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Person updated successfully",
	}, nil
}

func (o *Service) Delete(ctx context.Context, in *ms.EntityRequest) (*ms.Response, error) {
	if err := o.scyllaGlobal.ExecuteQuery("DELETE FROM people WHERE id = ?", in.Id); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Person deleted successfully",
	}, nil
}
