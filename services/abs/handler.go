package abs

import (
	"fmt"

	"github.com/reversTeam/go-ms/core"
)

type Handler struct{}

func NewHandler() core.GoMsHandlerInterface {
	return &Handler{}
}

func (o *Handler) RegisterHttp(gh *core.GoMsHttpServer, s core.GoMsServiceInterface) error {
	endpoint := fmt.Sprintf("%s:%d", gh.Grpc.Host, gh.Grpc.Port)
	return s.RegisterHttp(gh, endpoint)
}

func (o *Handler) RegisterGrpc(gs *core.GoMsGrpcServer, s core.GoMsServiceInterface) {
	s.RegisterGrpc(gs)
}
