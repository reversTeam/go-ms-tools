package child

import (
	"github.com/golang/protobuf/ptypes/empty"
	ms "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	"github.com/reversTeam/go-ms/core"
	pb "github.com/reversTeam/go-ms/services/child/protobuf"
	"github.com/reversTeam/go-ms/services/goms"
	"golang.org/x/net/context"
)

// Define the service structure
type Service struct {
	*goms.GoMsService
	config core.ServiceConfig
	pb.UnimplementedChildServer
}

// Instanciate the service without dependency because it's role of ServiceFactory
func NewService(name string, config core.ServiceConfig) *Service {
	s := &Service{
		GoMsService: goms.NewService(name, config),
	}

	return s
}

// This method is required for redister your service on the Http server
func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterChildHandlerFromEndpoint(gh.Ctx, gh.Mux, endpoint, gh.Grpc.Opts)
}

// This method is required for redister your service on the Grpc server
func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterChildServer(gs.Server, o)
}

// Endpoint :
//   - grpc : List
//   - http : Get /child
func (o *Service) List(ctx context.Context, in *empty.Empty) (*ms.Response, error) {
	return &ms.Response{
		Message: "Child List",
	}, nil
}

// Endpoint :
//   - grpc : Create
//   - http : POST /child
func (o *Service) Create(ctx context.Context, in *empty.Empty) (*ms.Response, error) {
	return &ms.Response{
		Message: "Child Create",
	}, nil
}

// Endpoint :
//   - grpc : Get
//   - http : GET /child/{id}
func (o *Service) Get(ctx context.Context, in *ms.EntityRequest) (*ms.Response, error) {
	return &ms.Response{
		Message: "Child View",
	}, nil
}

// Endpoint :
//   - grpc : Update
//   - http : PATCH /child/{id}
func (o *Service) Update(ctx context.Context, in *ms.EntityRequest) (*ms.Response, error) {
	return &ms.Response{
		Message: "Child Update",
	}, nil
}

// Endpoint :
//   - grpc : Delete
//   - http : PATCH /child/{id}
func (o *Service) Delete(ctx context.Context, in *ms.EntityRequest) (*ms.Response, error) {
	return &ms.Response{
		Message: "Child Delete",
	}, nil
}
