package handler

import (
	"context"
	pb "github.com/aaanger/event-analytics/grpc"
	"github.com/aaanger/event-analytics/internal/event/models"
	"github.com/aaanger/event-analytics/internal/event/repository"
)

type EventHandler struct {
	pb.UnimplementedEventServer
	repo repository.IEventRepository
}

func (h *EventHandler) SaveEvent(ctx context.Context, req *pb.EventRequest) (*pb.EventResponse, error) {
	err := h.repo.SaveEvent(ctx, models.Event{
		UserID:    req.UserId,
		PageURL:   req.Url,
		Timestamp: req.Timestamp,
	})
	if err != nil {
		return nil, err
	}

	return &pb.EventResponse{
		Status: "Event logged successfully",
	}, nil
}
