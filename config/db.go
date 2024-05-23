package config

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnect(ctx context.Context, dbConn string) (*pgxpool.Pool, error) {
	// Open connection
	var conn *pgxpool.Pool
	var err error
	for i := 0; i < 5; i++ {
		conn, err = pgxpool.New(ctx, dbConn)
		if err == nil {
			err = conn.Ping(ctx)
			if err == nil {
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Established a successful connection!")

	return conn, nil
}
