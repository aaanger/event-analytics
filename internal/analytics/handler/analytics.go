package handler

import (
	"context"
	pb "github.com/aaanger/event-analytics/grpc"
	"github.com/aaanger/event-analytics/internal/analytics/repository"
	"github.com/aaanger/event-analytics/pkg/lib"
)

type AnalyticsHandler struct {
	pb.UnimplementedAnalyticsServer
	repo repository.IAnalyticsRepository
}

func (h *AnalyticsHandler) GetUniqueUsers(ctx context.Context, req *pb.UniqueUsersRequest) (*pb.UniqueUsersResponse, error) {
	startDate, endDate, err := lib.ParseDates(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	users, err := h.repo.GetUniqueUsers(ctx, req.PageURL, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return &pb.UniqueUsersResponse{Users: users}, nil
}

func (h *AnalyticsHandler) GetPageViews(ctx context.Context, req *pb.PageViewsRequest) (*pb.PageViewsResponse, error) {
	startDate, endDate, err := lib.ParseDates(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	views, err := h.repo.GetPageViews(ctx, req.PageURL, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return &pb.PageViewsResponse{Views: views}, nil
}

func (h *AnalyticsHandler) GetTopPages(ctx context.Context, req *pb.TopPagesRequest) (*pb.TopPagesResponse, error) {
	startDate, endDate, err := lib.ParseDates(req.StartDate, req.EndDate)
	if err != nil {
		return nil, err
	}

	pages, err := h.repo.GetTopPages(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return &pb.TopPagesResponse{Pages: pages}, nil
}
