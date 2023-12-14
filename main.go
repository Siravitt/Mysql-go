package main

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Animal struct {
	Id   int
	Name string
}

// var db *sql.DB
var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", "root:@tcp(127.0.0.1:3306)/test_connection_go") // Connect to mysql database

	if err != nil {
		panic(err)
	}

	//! Query rows
	// animals, err := GetAnimals()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(animals)

	//! Query row
	// animal, err := GetAnimal(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(*animal)

	// ! Insert row
	// animal := Animal{
	// 	Id:   0,
	// 	Name: "Red Panda",
	// }
	// err = AddAnimal(animal)
	// if err != nil {
	// 	panic(err)
	// }

	//! Update row
	// animal := Animal{
	// 	Id:   1,
	// 	Name: "Rita",
	// }
	// err = UpdateAnimal(animal)
	// if err != nil {
	// 	panic(err)
	// }

	//! Delete row
	// animalId := 1
	// err = DeleteAnimal(animalId)
	// if err != nil {
	// 	panic(err)
	// }

	//! Query sqlX rows
	// animals, err := GetAnimalsX()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(animals)

	//! Query sqlX row
	// animal, err := GetAnimalByIdX(2)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(*animal)

	//! Delete sqlX row
	// animalId := 2
	// err = DeleteAnimalTran(animalId)
	// if err != nil {
	// 	panic(err)
	// }
}

func GetAnimalsX() ([]Animal, error) {
	query := "select id, name from animal"
	animals := []Animal{}
	err := db.Select(&animals, query)
	if err != nil {
		return nil, err
	}
	return animals, nil
}

func GetAnimalByIdX(id int) (*Animal, error) {
	query := "select id, name from animal where id = ?"
	animal := Animal{}
	err := db.Get(&animal, query, id)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

func GetAnimals() ([]Animal, error) {
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

func GetAnimal(id int) (*Animal, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := "select id, name from animal where id = ?"

	row := db.QueryRow(query, id)

	animal := Animal{}

	if err := row.Scan(&animal.Id, &animal.Name); err != nil {
		return nil, err
	}

	return &animal, nil
}

func AddAnimal(a Animal) error {
	query := "insert into animal (name) values (?)"

	result, err := db.Exec(query, a.Name)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected() // Check row affected
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("Cannot insert") // Create error
	}

	return nil
}

func UpdateAnimal(a Animal) error {
	query := "update animal set name = ? where id = ?"

	result, err := db.Exec(query, a.Name, a.Id)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("Cannot update")
	}

	return nil
}

func DeleteAnimal(id int) error {
	query := "delete from animal where id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("Cannot delete")
	}

	return nil
}

// ! Begin Transaction for rollback database
func DeleteAnimalTran(id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "delete from animal where id = ?"

	result, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if affect <= 0 {
		return errors.New("Cannot delete")
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
