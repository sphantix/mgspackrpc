package msgpackrpc

import (
	"io"
	"net"
	"net/rpc"

	"github.com/vmihailenco/msgpack/v5"
)

// Define custom MessagePack RPC client codec
type msgpackClientCodec struct {
	conn    io.ReadWriteCloser
	decoder *msgpack.Decoder
	encoder *msgpack.Encoder
}

func (c *msgpackClientCodec) WriteRequest(r *rpc.Request, body interface{}) error {
	if err := c.encoder.Encode(r); err != nil {
		return err
	}
	return c.encoder.Encode(body)
}

func (c *msgpackClientCodec) ReadResponseHeader(r *rpc.Response) error {
	return c.decoder.Decode(r)
}

func (c *msgpackClientCodec) ReadResponseBody(body interface{}) error {
	return c.decoder.Decode(body)
}

func (c *msgpackClientCodec) Close() error {
	return c.conn.Close()
}

// Function to create the custom client codec
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	return &msgpackClientCodec{
		conn:    conn,
		decoder: msgpack.NewDecoder(conn),
		encoder: msgpack.NewEncoder(conn),
	}
}

func NewClient(conn io.ReadWriteCloser) *rpc.Client {
	return rpc.NewClientWithCodec(NewClientCodec(conn))
}

// Function to dial the RPC server using MessagePack codec
func Dial(network, address string) (*rpc.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	// Use the custom MessagePack client codec
	return rpc.NewClientWithCodec(NewClientCodec(conn)), nil
}
