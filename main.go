package main

import (
	"net"

	pb "github.com/chat-rpc/ggrpc-svr-demo/pb"

	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port	= ":50051"
)

type server struct {

}

func (s *server) SayHello(ctx context.Context,in *pb.HelloRequest) (*pb.HelloReply,error){
	return &pb.HelloReply{Message:"Hello " + in.Name},nil
}

func main(){
	lis,err := net.Listen("tcp",port);


	if err != nil {
		fmt.Print("failed to listen : %v",err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})

	fmt.Print("listen on ",port);

	s.Serve(lis);
}
