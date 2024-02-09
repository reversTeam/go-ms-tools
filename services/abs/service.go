package abs

import (
	// "fmt"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	"github.com/reversTeam/go-ms/core"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"golang.org/x/net/context"
)

// Define the service structure
type Service struct {
	Name    string
	Ctx     *core.Context
	config  core.ServiceConfig
	Handler core.GoMsHandlerInterface
	Metrics *Metrics
	pb.UnimplementedAbsServer
}

// Instanciate the service without dependency because it's role of ServiceFactory
func NewService(ctx *core.Context, name string, config core.ServiceConfig) *Service {
	o := &Service{
		Ctx:     ctx,
		Name:    name,
		config:  config,
		Handler: NewHandler(),
		Metrics: NewMetrics(name),
	}
	return o
}

// TODO: Need to encapsule request with it
func (o *Service) Log(message string) {
	ctx := context.Background()
	tracer := otel.Tracer(o.Name)
	_, span := tracer.Start(ctx, message)
	defer span.End()

	span.SetAttributes(attribute.String("operation", "list"))
	span.SetStatus(codes.Ok, "Successful operation")
}

// Get the service name
func (o *Service) GetName() string {
	return o.Name
}

// Get Handler describe in the interface
func (o *Service) GetHandler() core.GoMsHandlerInterface {
	return o.Handler
}

// This method is required for redister your service on the Http server
func (o *Service) RegisterHttp(gh *core.GoMsHttpServer, endpoint string) error {
	return pb.RegisterAbsHandlerFromEndpoint(gh.Ctx.Main, gh.Mux, endpoint, gh.Grpc.Opts)
}

// This method is required for redister your service on the Grpc server
func (o *Service) RegisterGrpc(gs *core.GoMsGrpcServer) {
	pb.RegisterAbsServer(gs.Server, o)
}

// Endpoint :
//   - grpc : List
//   - http : Get /abs
func (o *Service) List(ctx context.Context, in *empty.Empty) (*pb.Response, error) {
	return &pb.Response{
		Message: "Abs List",
	}, nil
}

// Endpoint :
//   - grpc : Create
//   - http : POST /abs
func (o *Service) Create(ctx context.Context, in *empty.Empty) (*pb.Response, error) {
	return &pb.Response{
		Message: "Abs Create",
	}, nil
}

// Endpoint :
//   - grpc : Get
//   - http : GET /abs/{id}
func (o *Service) Get(ctx context.Context, in *pb.EntityRequest) (*pb.Response, error) {
	return &pb.Response{
		Message: "Abs View",
	}, nil
}

// Endpoint :
//   - grpc : Update
//   - http : PATCH /abs/{id}
func (o *Service) Update(ctx context.Context, in *pb.EntityRequest) (*pb.Response, error) {
	return &pb.Response{
		Message: "Abs Update",
	}, nil
}

// Endpoint :
//   - grpc : Delete
//   - http : DELETE /abs/{id}
func (o *Service) Delete(ctx context.Context, in *pb.EntityRequest) (*pb.Response, error) {
	return &pb.Response{
		Message: "Abs Delete",
	}, nil
}
