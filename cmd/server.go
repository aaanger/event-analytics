package main

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	analyticsGRPC "github.com/aaanger/event-analytics/internal/analytics/handler"
	eventGRPC "github.com/aaanger/event-analytics/internal/event/handler"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	engine *grpc.Server
	port   string
	db     clickhouse.Conn
}

func NewServer(db clickhouse.Conn, port string) *Server {
	grpcServer := grpc.NewServer()

	return &Server{
		engine: grpcServer,
		port:   port,
		db:     db,
	}
}

func (s Server) Run() error {
	eventGRPC.Register(s.engine, s.db)
	analyticsGRPC.Register(s.engine, s.db)

	lis, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}

	err = s.engine.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
