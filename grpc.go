package grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type grpcServer struct {
	server *grpc.Server
}

func NewGrpc() *grpcServer {
	srv := grpc.NewServer()

	return &grpcServer{
		server: srv,
	}
}

type RegisterFunc func(*grpc.Server)

// register server to grpc
func (g *grpcServer) Register(servers ...RegisterFunc) *grpcServer {

	for _, registerServer := range servers {
		registerServer(g.server)
	}

	return g
}

func (g *grpcServer) Run(addr string) {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic("err")
	}

	fmt.Println("Listening and serving Rpc on ", addr)

	if err := g.server.Serve(lis); err != nil {
		panic(err)
	}

}
