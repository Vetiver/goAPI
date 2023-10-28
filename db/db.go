package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Data struct {
	name string
    id   int
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

func Insert(name string) string {
	pool := DbStart()

	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Println(fmt.Errorf("unable to acquire a database connection: %v", err))
		return "не коннектится пупсик"
	}

	row := conn.QueryRow(context.Background(),
		"INSERT INTO test(name) VALUES ($1) RETURNING id;", name)
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

func DeliteById(id int) string {
	pool := DbStart()

	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Println(fmt.Errorf("unable to acquire a database connection: %v", err))
		return "ошибка соединения"
	}

	row := conn.QueryRow(context.Background(),
		"DELETE FROM test WHERE id=$1 RETURNING id;", id)
	//после коннекта прописываем запрос на DELETE и возвращаем id
	err = row.Scan(&id)
	if err != nil {
		fmt.Println(fmt.Errorf("unable to DELITE: %v", err))
		//если ты тупой, то тебе вернет ошибку пупсик
		return "неправильный запрос, карточки с таким id не существует"
	}
	return "успешное удаление"
}

func GetAllNames() ([]Data, error) {
	pool := DbStart()
	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
	 return nil, fmt.Errorf("unable to acquire a database connection: %v", err)
	}
   
	rows, err := conn.Query(context.Background(),
	 "SELECT id, name FROM test")
	if err != nil {
	 return nil, fmt.Errorf("unable to retrieve data from database: %v", err)
	}
	defer rows.Close()
	var data []Data
	for rows.Next() {
	 var d Data
	 err = rows.Scan(&d.id, &d.name)
	 if err != nil {
	  return nil, fmt.Errorf("unable to scan row: %v", err)
	 }
	 data = append(data, d)
	 fmt.Println(fmt.Errorf("тут: %+v", data))
	}
   
   
	return data, nil
   }

func UpdateName(name string, id int) any {
	pool := DbStart()

	conn, err := pool.Acquire(context.Background())
	//Acqure - забирает одно соединение с бд из pool
	if err != nil {
		fmt.Println(fmt.Errorf("unable to acquire a database connection: %v", err))
		return "ошибка соединения"
	}

	row := conn.QueryRow(context.Background(),
		"UPDATE test SET name = ? WHERE id = ?", name, id)
	//после коннекта прописываем запрос на DELETE и возвращаем id
	//интициализируем переменную id
	row.Scan()
	//сканируем значение id
	if err != nil {
		fmt.Println(fmt.Errorf("unable to INSERT: %v", err))
		//если ты тупой, то тебе вернет ошибку пупсик
		return "неправильный запрос придурок или ты пытаешься вогнать не тот тип данных(осел)"
	}
	return "успешный апдейт"
}
