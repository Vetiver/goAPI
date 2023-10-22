package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Record struct {
	Id int `json:"id"`
}

func DbStart() *pgxpool.Pool {
	urlExample := "postgres://postgres:228@localhost:5432/postgres"
	dbpool, err := pgxpool.New(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v", err)
		os.Exit(1)
	}

	return dbpool
}

func Insert() string {
	pool := DbStart()

	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Println(fmt.Errorf("unable to acquire a database connection: %v", err))
		return "не коннектится пупсик"
	}

	row := conn.QueryRow(context.Background(),
		"INSERT INTO test(name) VALUES ($1) RETURNING id;", "rrrr")
	//после коннекта прописываем запрос на Insert и возвращаем значение id
	var id uint64
	//интициализируем переменную id
	err = row.Scan(&id)
	//сканируем значение id
	if err != nil {
		fmt.Println(fmt.Errorf("unable to INSERT: %v", err))
		//если ты тупой, то тебе вернет ошибку пупсик
		return "тупорылая ты ослица, что ты пытался сделать?"
	}
	return "успешное добавление"
}


func DeliteById() string {
	pool := DbStart()

	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Println(fmt.Errorf("unable to acquire a database connection: %v", err))
		return "ошибка соединения"
	}

	row := conn.QueryRow(context.Background(),
		"DELETE FROM test WHERE id=$1 RETURNING id;", 1)
	//после коннекта прописываем запрос на DELETE и возвращаем id
	var id uint64
	//интициализируем переменную id
	err = row.Scan(&id)
	//сканируем значение id
	if err != nil {
		fmt.Println(fmt.Errorf("unable to INSERT: %v", err))
		//если ты тупой, то тебе вернет ошибку пупсик
		return "неправильный запрос придурок или ты пытаешься вогнать не тот тип данных(осел)"
	}
	return "успешное удаление"
}


func GetAllNames() any {
	pool := DbStart()

	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Println(fmt.Errorf("unable to acquire a database connection: %v", err))
		return "ошибка соединения"
	}

	row := conn.QueryRow(context.Background(),
		"SELECT * FROM test")
	//после коннекта прописываем запрос на получение инфы о всех таблицах
	err = row.Scan(&row)
	users, err := json.Marshal(err)
	//сканируем значение id
	if err != nil {
		fmt.Println(fmt.Errorf("unable to INSERT: %v", err))
		//если ты тупой, то тебе вернет ошибку пупсик
		return "неправильный запрос придурок или ты пытаешься вогнать не тот тип данных(осел)"
	}
	return users
}