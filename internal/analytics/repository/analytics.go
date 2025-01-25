package repository

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/aaanger/event-analytics/grpc"
	"time"
)

type IAnalyticsRepository interface {
	GetUniqueUsers(ctx context.Context, url string, startDate, endDate time.Time) (int64, error)
	GetPageViews(ctx context.Context, url string, startDate, endDate time.Time) (int64, error)
	GetTopPages(ctx context.Context, startDate, endDate time.Time) ([]*grpc.PageStats, error)
}

type AnalyticsRepository struct {
	db clickhouse.Conn
}

func NewAnalyticsRepository(db clickhouse.Conn) *AnalyticsRepository {
	return &AnalyticsRepository{
		db: db,
	}
}

func (r *AnalyticsRepository) GetUniqueUsers(ctx context.Context, url string, startDate, endDate time.Time) (int64, error) {
	var users int64

	row := r.db.QueryRow(ctx, `SELECT COUNT(DISTINCT user_id) FROM events WHERE page_url = ? AND timestamp BETWEEN ? AND ?`, url, startDate, endDate)
	err := row.Scan(&users)
	if err != nil {
		return 0, err
	}

	return users, nil
}

func (r *AnalyticsRepository) GetPageViews(ctx context.Context, url string, startDate, endDate time.Time) (int64, error) {
	var views int64

	row := r.db.QueryRow(ctx, `SELECT SUM(views) FROM page_views WHERE page_url = ? AND timestamp BETWEEN ? AND ?`, url, startDate, endDate)
	err := row.Scan(&views)
	if err != nil {
		return 0, err
	}

	return views, nil
}

func (r *AnalyticsRepository) GetTopPages(ctx context.Context, startDate, endDate time.Time) ([]*grpc.PageStats, error) {
	rows, err := r.db.Query(ctx, `SELECT (page_url, views) FROM events WHERE timestamp BETWEEN ? AND ? GROUP BY page_url ORDER BY views`)
	if err != nil {
		return nil, err
	}

	var pages []*grpc.PageStats

	for rows.Next() {
		var page grpc.PageStats

		err = rows.Scan(&page.PageURL, &page.Views)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &page)
	}

	return pages, rows.Err()
}
