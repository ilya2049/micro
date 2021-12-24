package grpcutil

import (
	"context"

	"google.golang.org/grpc"
)

func NewConnection(
	ctx context.Context,
	clientConnection *grpc.ClientConn,
	closeFunc func(),
) *Connection {
	return &Connection{
		clientConnection: clientConnection,
		ctx:              ctx,
		closeFunc:        closeFunc,
	}
}

type Connection struct {
	clientConnection *grpc.ClientConn

	ctx       context.Context
	closeFunc func()
}

func (c *Connection) ClientConnection() *grpc.ClientConn {
	return c.clientConnection
}

func (c *Connection) Context() context.Context {
	return c.ctx
}

func (c *Connection) Close() {
	c.closeFunc()
}
