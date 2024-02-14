package abs

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/reversTeam/go-ms-tools/services/abs/protobuf"
	"github.com/reversTeam/go-ms/core"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Define the service structure
type Service struct {
	Name          string
	Ctx           *core.Context
	config        core.ServiceConfig
	Handler       core.GoMsHandlerInterface
	Metrics       *Metrics
	Client        pb.AbsClient
	ClientManager *core.GrpcClientManager
	pb.UnimplementedAbsServer
}

// Instanciate the service without dependency because it's role of ServiceFactory
func NewService(ctx *core.Context, name string, config core.ServiceConfig) *Service {
	o := &Service{
		Ctx:           ctx,
		Name:          name,
		config:        config,
		Handler:       NewHandler(),
		Metrics:       NewMetrics(name),
		ClientManager: nil,
		Client:        nil,
	}
	return o
}

func (o *Service) SetClientManager(clientManager *core.GrpcClientManager) {
	o.ClientManager = clientManager
}

func (o *Service) GetMiddlewaresConf() map[string][]string {
	result := make(map[string][]string)

	middlewares, ok := o.config.Config["middlewares"]
	if !ok {
		return result
	}

	middlewaresMap, ok := middlewares.(map[string]interface{})
	if !ok {
		return result
	}

	for key, value := range middlewaresMap {
		valueSlice, ok := value.([]interface{})
		if !ok {
			continue
		}

		stringSlice := make([]string, len(valueSlice))
		for i, v := range valueSlice {
			vStr, ok := v.(string)
			if !ok {
				continue
			}
			stringSlice[i] = vStr
		}

		result[key] = stringSlice
	}

	return result
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

func (o *Service) GetClient(conn *grpc.ClientConn) any {
	return pb.NewAbsClient(conn)
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
