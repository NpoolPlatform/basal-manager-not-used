package api

import (
	api1 "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	api1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	api1.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
