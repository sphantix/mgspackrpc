# msgpackrpc
golang msgpack codec compatible with net/rpc

## Usage
### Server
```go
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
```

### Client
```go
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
```

## Examples
code above can be found in [examples](./examples)

```
go run examples/server/server.go
go run examples/client/client.go
```