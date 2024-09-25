package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/sphantix/msgpackrpc"
)

type APIService struct{}

func (s *APIService) Hello(args string, reply *string) error {
	*reply = "Hello, " + args
	return nil
}

type AddArgs struct {
	A, B int
}

type AddReply struct {
	Result int
}

func (s *APIService) Add(args AddArgs, reply *AddReply) error {
	reply.Result = args.A + args.B
	return nil
}

func main() {
	api := new(APIService)
	rpc.Register(api)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeCodec(msgpackrpc.NewServerCodec(conn))
	}
}
