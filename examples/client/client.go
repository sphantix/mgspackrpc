package main

import (
	"fmt"
	"log"
	"net"

	"github.com/sphantix/msgpackrpc"
)

type AddArgs struct {
	A, B int
}

type AddReply struct {
	Result int
}

func main() {
	// Connect to the server using TCP
	conn, err := net.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	client := msgpackrpc.NewClient(conn)
	defer client.Close()

	// Call the RPC method "APIService.Hello"
	var reply string
	var args = "World"

	err = client.Call("APIService.Hello", args, &reply)
	if err != nil {
		log.Fatalf("RPC call failed: %v", err)
	}

	// Print the result
	fmt.Printf("Result of APIService.Hello: %s\n", reply) // Expected output: 15

	var addArgs = AddArgs{10, 5}
	var addReply AddReply
	err = client.Call("APIService.Add", addArgs, &addReply)
	if err != nil {
		log.Fatalf("RPC call failed: %v", err)
	}
	fmt.Printf("Result of APIService.Add: %d\n", addReply.Result) // Expected output: 15
}
