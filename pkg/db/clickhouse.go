package db

import (
	"context"
	"crypto/tls"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type ClickHouseConfig struct {
	Username string
	Password string
	Database string
}

func Open(cfg ClickHouseConfig) (driver.Conn, error) {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"<CLICKHOUSE_SECURE_NATIVE_HOSTNAME>:9440"},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.Username,
			Password: cfg.Password,
		},
		ClientInfo: clickhouse.ClientInfo{
			Products: []struct {
				Name    string
				Version string
			}{
				{
					Name:    "event-analytics",
					Version: "0.1",
				},
			},
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
	})

	if err != nil {
		return nil, err
	}

	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, err
}
