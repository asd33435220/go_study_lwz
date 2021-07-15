package main

import (
	pb "../pb"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct {
}

func (s *server) SayHello(c context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main(){
	lis ,err := net.Listen("tcp",":8972")
	if err != nil {
		fmt.Printf("failed to listen: %v",err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to save %v",err)
		return
	}
}
