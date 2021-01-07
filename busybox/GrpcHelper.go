package busybox

import (
	grpc "google.golang.org/grpc"
	"log"
)

type GrpcHelper struct {
}

func (gh GrpcHelper) Conn(host string) *grpc.ClientConn {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}

	return conn
}
