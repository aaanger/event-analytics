package handler

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	pb "github.com/aaanger/event-analytics/grpc"
	"github.com/aaanger/event-analytics/internal/event/repository"
	"google.golang.org/grpc"
)

func Register(grpcServer *grpc.Server, db clickhouse.Conn) {
	repo := repository.NewEventRepository(db)
	pb.RegisterEventServer(grpcServer, &EventHandler{repo: repo})
}
