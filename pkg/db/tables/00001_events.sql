-- +goose Up
-- +goose StatementBegin
CREATE TABLE events (
    user_id UInt64,
    page_url String,
    timestamp DateTime,
    views UInt64 DEFAULT 1
) ENGINE = SummingMergeTree()
ORDER BY (page_url, timestamp)
PARTITION BY toYYYYMM(timestamp);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE events;
-- +goose StatementEnd
