package v1

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ControllerServer struct{}

var _ ControllerServiceServer = (*ControllerServer)(nil)

func Serve(port int) error {
	grpcServer := grpc.NewServer()
	var server ControllerServer
	RegisterControllerServiceServer(grpcServer, &server)

	fmt.Printf("Serving on *:%d\n", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("unable to start tcp server: %w", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("unable to start grpc server: %w", err)
	}

	return nil
}

//nolint:revive
func (c ControllerServer) RegisterApp(ctx context.Context, request *RegisterAppRequest) (*emptypb.Empty, error) {
	fmt.Printf("Registering app: %s\n", request.Id)

	return &emptypb.Empty{}, nil
}
