package api

import (
	"context"

	basal "github.com/NpoolPlatform/message/npool/basal/gw/v1"

	api1 "github.com/NpoolPlatform/basal-manager/api/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	basal.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	basal.RegisterManagerServer(server, &Server{})
	api1.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
