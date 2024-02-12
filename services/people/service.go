package people

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reversTeam/go-ms-tools/pkg/scylla"
	"github.com/reversTeam/go-ms-tools/services/abs"
	ms "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	pb "github.com/reversTeam/go-ms-tools/services/people/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/net/context"
)

// Define the service structure
type Service struct {
	*abs.Service
	config core.ServiceConfig
	pb.UnimplementedPeopleServer
	scylla *scylla.ScyllaDBManager
}

func NewService(ctx *core.Context, name string, config core.ServiceConfig) *Service {
	scyllaDBManager, err := scylla.NewScyllaDBManager()
	if err != nil {
		panic(err)
	}

	return &Service{
		Service: abs.NewService(ctx, name, config),
		scylla:  scyllaDBManager,
	}
}

func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterPeopleHandlerFromEndpoint(gh.Ctx.Main, gh.Mux, endpoint, gh.Grpc.Opts)
}

func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterPeopleServer(gs.Server, o)
}

func (o *Service) List(in *empty.Empty, stream pb.People_ListServer) error {
	_, scyllaSpan := core.Trace(stream.Context(), "people", "List")
	defer scyllaSpan.End()

	var peoples []pb.PeopleResponse // Définissez PeopleResponse en Go pour correspondre au message protobuf
	err := o.scylla.Get(&peoples, "SELECT id, firstname, lastname, email, birthday FROM glb.people")
	if err != nil {
		return err
	}

	for _, people := range peoples {
		resp := &pb.PeopleResponse{
			Id:        people.Id,
			Firstname: people.Firstname,
			Lastname:  people.Lastname,
			Email:     people.Email,
			Birthday:  people.Birthday,
		}
		if err := stream.Send(resp); err != nil {
			return err
		}
	}

	return nil
}

func (o *Service) Get(ctx context.Context, in *ms.EntityRequest) (*pb.PeopleResponse, error) {
	id, err := gocql.ParseUUID(in.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	ctx, span := core.Trace(ctx, "people", "Get")
	defer span.End()

	var people pb.PeopleResponse
	err = o.scylla.GetOne(&people, "SELECT id, firstname, lastname, email, birthday FROM glb.people WHERE id = ?", id)

	return &people, err
}

func (o *Service) Create(ctx context.Context, in *pb.PeopleCreateParams) (*ms.Response, error) {
	id := gocql.TimeUUID()
	if err := o.scylla.ExecuteQuery("INSERT INTO glb.people (id, firstname, lastname, email, birthday) VALUES (?, ?, ?, ?, ?)",
		id, in.Firstname, in.Lastname, in.Email, in.Birthday); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Person created successfully",
	}, nil
}

func (o *Service) Update(ctx context.Context, in *pb.PeopleUpdateParams) (*ms.Response, error) {
	if err := o.scylla.ExecuteQuery("UPDATE glb.people SET firstname = ?, lastname = ?, email = ?, birthday = ? WHERE id = ?",
		in.Firstname, in.Lastname, in.Email, in.Birthday, in.Id); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Person updated successfully",
	}, nil
}

func (o *Service) Delete(ctx context.Context, in *ms.EntityRequest) (*ms.Response, error) {
	if err := o.scylla.ExecuteQuery("DELETE FROM glb.people WHERE id = ?", in.Id); err != nil {
		return nil, err
	}

	return &ms.Response{
		Message: "Person deleted successfully",
	}, nil
}
