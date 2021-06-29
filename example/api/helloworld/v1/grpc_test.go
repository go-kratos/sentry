package v1

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"google.golang.org/grpc"
	"testing"
)

func TestGRPC(t *testing.T)  {
	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := NewGreeterClient(conn)
	resp, err := c.SayHello(context.TODO(), &HelloRequest{Name:"grpc"}, )
	if err != nil {
		panic(err)
	}
	spew.Dump(resp)
}