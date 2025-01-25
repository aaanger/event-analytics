package repository

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2"

	"github.com/aaanger/event-analytics/internal/event/models"
)

type IEventRepository interface {
	SaveEvent(ctx context.Context, event models.Event) error
}

type EventRepository struct {
	db clickhouse.Conn
}

func NewEventRepository(db clickhouse.Conn) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) SaveEvent(ctx context.Context, event models.Event) error {
	err := r.db.Exec(ctx, `INSERT INTO events (user_id, page_url, timestamp, views) VALUES(?, ?, ?, 1)`,
		event.UserID, event.PageURL, event.Timestamp)
	if err != nil {
		return err
	}

	return nil
}
