package grpcClient

import (
	pb "github.com/FakJeongTeeNhoi/api-gateway/generated/space"

	"google.golang.org/grpc"

	"os"
)

type spaceClient struct {
	Client pb.SpaceServiceClient
}

var SpaceClient *spaceClient

// NewClient creates a new grpcClient client
func NewClient() (*grpc.ClientConn, error) {
	address := os.Getenv("GRPC_URL")

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewSpaceClient() {
	conn, err := NewClient()
	if err != nil {
		panic(err)
	}

	client := pb.NewSpaceServiceClient(conn)

	SpaceClient = &spaceClient{
		Client: client,
	}
}
