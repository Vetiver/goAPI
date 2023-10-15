package handlers

import "github.com/gin-gonic/gin"
import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"
)
type Record struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
  }

func Any(c *gin.Context) {
	c.JSON(200, )
}

func Insert() string {
	urlExample := "postgres://postgres:228@localHost:5432/postgres"
	var rec Record
	err := json.NewDecoder().Decode(&rec)
	if err != nil { // bad request
	  return ""
	}
  
	conn, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Errorf("Unable to acquire a database connection: %v\n", err)
	  return ""
	}
  
	row := conn.QueryRow(context.Background(),
	  "INSERT INTO phonebook (name, phone) VALUES ($1, $2) RETURNING id",
	  rec.Name, rec.Phone)
	var id uint64
	err = row.Scan(&id)
	if err != nil {
		fmt.Errorf("Unable to INSERT: %v\n", err)

	  return ""
	}
  
	resp := make(map[string]string, 1)
	resp["id"] = strconv.FormatUint(id, 10)
	err = json.NewEncoder().Encode(resp)
	if err != nil {
		fmt.Errorf("Unable to encode json: %v\n", err)

	  return ""
	}
	return "что-то"
  }