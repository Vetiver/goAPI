package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Data struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type DB struct {
	pool *pgxpool.Pool
}

func NewDB(pool *pgxpool.Pool) *DB {
	return &DB{
		pool: pool,
	}
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

func (db DB) Insert(name Data) (*Data, error) {


    conn, err := db.pool.Acquire(context.Background())
    if err != nil {
        return nil, fmt.Errorf("unable to acquire a database connection: %v", err)
    }
	defer conn.Release()
    err = conn.QueryRow(context.Background(),
        "INSERT INTO test(name) VALUES ($1) RETURNING id", name.Name).Scan(&name.Id)
    if err != nil {
        return nil, fmt.Errorf("unable to INSERT: %v", err)
    }

    return &name, nil
}

func (db DB) DeleteById(id int) error {

	conn, err := db.pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		return fmt.Errorf("unable to acquire a database connection: %v", err)
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(),
		"DELETE FROM test WHERE id=$1 RETURNING id;", id)
	//после коннекта прописываем запрос на DELETE и возвращаем id
	err = row.Scan(&id)
	if err != nil {
		//если ты тупой, то тебе вернет ошибку пупсик
		return fmt.Errorf("unable to DELITE: %v", err)
	}
	return nil
}

func (db DB) GetAllNames() ([]Data, error) {
	conn, err := db.pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		return nil, fmt.Errorf("unable to acquire a database connection: %v", err)
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(),
		"SELECT id, name FROM test")
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve data from database: %v", err)
	}
	defer rows.Close()
	var data []Data
	for rows.Next() {
		var d Data
		err = rows.Scan(&d.Id, &d.Name)
		if err != nil {
			return nil, fmt.Errorf("unable to scan row: %v", err)
		}
		data = append(data, d)
		fmt.Println(fmt.Errorf("тут: %+v", data))
	}
	return data, nil
}

func (db DB) UpdateName(name Data) (*Data, error) {

	conn, err := db.pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		return nil, fmt.Errorf("unable to acquire a database connection: %v", err)
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(),
		"UPDATE test SET name = $1  WHERE id = $2 ", name.Name, name.Id)
	//после коннекта прописываем запрос на DELETE и возвращаем id
	//интициализируем переменную id
	row.Scan()
	//сканируем значение id
	if err != nil {
		//если ты тупой, то тебе вернет ошибку пупсик
		return nil, fmt.Errorf("unable to UPDATE: %v", err)
	}
	return &name, nil
}
