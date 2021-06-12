package main

import (
	"context"
	"github.com/code-newbee/protocol/geeker"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:50001"
)

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil{
		log.Fatal("dial svc error")
	}

	defer conn.Close()

	c := geeker.NewGeekerClient(conn)

	rsp, err := c.SayHello(context.Background(), &geeker.SayHelloRequest{RequestId: "hello in client"})
	if err != nil{
		log.Fatal(err)
	}
	log.Println(rsp)
}