package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	gomicro "go-micro/proto"
)

type GoMicro struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GoMicro) Call(ctx context.Context, req *gomicro.Request, rsp *gomicro.Response) error {
	log.Info("Received GoMicro.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *GoMicro) Stream(ctx context.Context, req *gomicro.StreamingRequest, stream gomicro.GoMicro_StreamStream) error {
	log.Infof("Received GoMicro.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&gomicro.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *GoMicro) PingPong(ctx context.Context, stream gomicro.GoMicro_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&gomicro.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
