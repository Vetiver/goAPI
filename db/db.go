package db

import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Record struct {
	Id    int    `json:"id"`
  }

func DbStart() *pgxpool.Pool {
	urlExample := "postgres://postgres:228@localhost:5432/postgres"
	dbpool, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	return dbpool
}

func Insert() string {
	pool := DbStart()
  
	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Errorf("Unable to acquire a database connection: %v\n", err)
	  return ""
	}
  
	row := conn.QueryRow(context.Background(),
	  "INSERT INTO test(name) VALUES ($1);", "gay")
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		fmt.Errorf("Unable to INSERT: %v\n", err)

	  return ""
	}
	return "что-то"
  }
