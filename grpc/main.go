package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	service "github.com/walker2/k8s-goapm-examples/grpc/protobuf"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Send(ctx context.Context, req *service.Request) (*service.Response, error) {
	log.Printf("Send function was invoked %v\n", req)
	milli := rand.Int63n(5000)
	log.Printf("Lifting weights, %d millis ( ͡° ͜ʖ ͡°)", milli)
	time.Sleep(time.Duration(milli) * time.Millisecond)

	if ctx.Err() == context.Canceled {
		log.Printf("Client canceled, stopping\n")
		return nil, status.Error(codes.Canceled, "Client canceled, stopping")
	}

	log.Printf("Done ᕙ( ͡° ͜ʖ ͡°)ᕗ")

	return &service.Response{}, nil
}

func startServer() {
	fmt.Println("gRPC server is listening on port 50051")
	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	// Important, apm interceptor
	opts = append(opts, grpc.UnaryInterceptor(apmgrpc.NewUnaryServerInterceptor()))

	s := grpc.NewServer(opts...)
	service.RegisterServiceAServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func startClient() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(apmgrpc.NewUnaryClientInterceptor()))

	conn, err := grpc.Dial(":50051", opts...)
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	defer conn.Close()

	c := service.NewServiceAClient(conn)

	for {
		callWithDeadline(c, 4000)
	}

}

func callWithDeadline(c service.ServiceAClient, timeout time.Duration) {
	log.Println("Calling server")
	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout*time.Millisecond,
	)
	defer cancel()

	res, err := c.Send(ctx, &service.Request{})
	if err != nil {

		statusErr, ok := status.FromError(err)

		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Printf("Timeout was hit, deadline exceeded")
			} else {
				log.Fatalf("Unexpected error %v\n", err)
			}
			return
		}
		log.Fatalf("Error while calling server %v", err)
	}
	log.Printf("res: %v\n", res)

	time.Sleep(time.Duration(5000) * time.Millisecond)
}

func main() {
	go startClient()

	startServer()
}
