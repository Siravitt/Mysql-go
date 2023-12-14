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

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test_connection_go")

	if err != nil {
		panic(err)
	}

	animals, err := GetAnimals()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(animals)
}

func GetAnimals() ([]Animal, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := "select id, name from animal"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	animals := []Animal{}

	for rows.Next() {
		animal := Animal{}
		err = rows.Scan(&animal.Id, &animal.Name)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}
