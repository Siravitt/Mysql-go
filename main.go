package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Animal struct {
	Id   int
	Name string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test_connection_go")

	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	query := "select id, name from animal"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	animals := []Animal{}

	for rows.Next() {
		animal := Animal{}
		err = rows.Scan(&animal.Id, &animal.Name)
		if err != nil {
			panic(err)
		}
		animals = append(animals, animal)
	}

	fmt.Println(animals)
}
