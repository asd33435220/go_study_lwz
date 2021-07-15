package main

import (
	pb "../pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{
		Name: "zhuyuchen",
	})
	if err != nil {
		fmt.Printf("could not greet : %v",err)
	}
	fmt.Printf("Greeting: %s !\n",r.Message)

}
