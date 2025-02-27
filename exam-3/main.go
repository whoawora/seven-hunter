package main

import (
	"context"
	"log"
	"net"
	"pie-fire-dire/handler"
	"pie-fire-dire/proto"
	"pie-fire-dire/service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type beefServer struct {
	proto.UnimplementedBeefServiceServer
}

func (s *beefServer) GetBeefSummary(ctx context.Context, _ *proto.Empty) (*proto.BeefSummaryResponse, error) {
	counts, err := service.FetchAndCountBeef()
	if err != nil {
		return nil, err
	}
	resp := &proto.BeefSummaryResponse{Beef: make(map[string]int32)}
	for k, v := range counts {
		resp.Beef[k] = int32(v)
	}
	return resp, nil
}

func main() {
	go func() {
		r := gin.Default()
		r.GET("/beef/summary", handler.GetBeefSummary)
		log.Fatal(r.Run(":8000"))
	}()

	lis, err := net.Listen("tcp", ":18000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterBeefServiceServer(grpcServer, &beefServer{})
	log.Printf("GRPC server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
