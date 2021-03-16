package busybox

import (
	"crypto/tls"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

func (gh GrpcHelper) ConnWithSSL(host string) *grpc.ClientConn {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(credentials.NewTLS(config)))

	if err != nil {
		log.Fatalf("Connect failed: %v", err)
	}

	return conn
}
