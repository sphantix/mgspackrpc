package msgpackrpc

import (
	"io"
	"net/rpc"

	"github.com/vmihailenco/msgpack/v5"
)

// Define custom MessagePack RPC server codec
type msgpackServerCodec struct {
	conn    io.ReadWriteCloser
	decoder *msgpack.Decoder
	encoder *msgpack.Encoder
}

func (c *msgpackServerCodec) ReadRequestHeader(r *rpc.Request) error {
	return c.decoder.Decode(r)
}

func (c *msgpackServerCodec) ReadRequestBody(body interface{}) error {
	return c.decoder.Decode(body)
}

func (c *msgpackServerCodec) WriteResponse(r *rpc.Response, body interface{}) error {
	if err := c.encoder.Encode(r); err != nil {
		return err
	}
	if err := c.encoder.Encode(body); err != nil {
		return err
	}
	return nil
}

func (c *msgpackServerCodec) Close() error {
	return c.conn.Close()
}

func NewServerCodec(conn io.ReadWriteCloser) rpc.ServerCodec {
	return &msgpackServerCodec{
		conn:    conn,
		decoder: msgpack.NewDecoder(conn),
		encoder: msgpack.NewEncoder(conn),
	}
}

func ServeConn(conn io.ReadWriteCloser) {
	rpc.ServeCodec(NewServerCodec(conn))
}
