package grpc

import (
	"mytro-backend-content/internal/infrastructure/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewContextClient(addr string) (pb.ContentStorageClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewContentStorageClient(conn), nil
}
