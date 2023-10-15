package db

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Record struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
  }

func DbStart() {
	urlExample := "postgres://postgres:228@localHost:5432/postgres"
	dbpool, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	

	fmt.Println(greeting)
}

func Insert(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	var rec Record
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil { // bad request
	  w.WriteHeader(400)
	  return
	}
  
	conn, err := p.Acquire(context.Background())
	if err != nil {
		fmt.Errorf("Unable to acquire a database connection: %v\n", err)
	  w.WriteHeader(500)
	  return
	}
	defer conn.Release()
  
	row := conn.QueryRow(context.Background(),
	  "INSERT INTO phonebook (name, phone) VALUES ($1, $2) RETURNING id",
	  rec.Name, rec.Phone)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		fmt.Errorf("Unable to INSERT: %v\n", err)
	  w.WriteHeader(500)
	  return
	}
  
	resp := make(map[string]string, 1)
	resp["id"] = strconv.FormatUint(id, 10)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		fmt.Errorf("Unable to encode json: %v\n", err)
	  w.WriteHeader(500)
	  return
	}
  }