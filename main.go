package main

import (
	"context"
	"github.com/code-newbee/protocol/geeker"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "192.168.0.105:50051"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		log.Fatal("dial svc error")
	}

	defer conn.Close()

	c := geeker.NewGreekerClient(conn)

	rsp, err := c.SayHello(context.Background(), &geeker.HelloRequest{Name: "hello in client go-gRpc"})
	if err != nil{
		log.Fatal(err)
	}
	log.Println(rsp)
}