package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Customers struct {
	Id        string
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	BirtDate  time.Time `db:"birth_date"`
	Address   sql.NullString
	Status    int
	Username  sql.NullString
	Password  sql.NullString
	Email     sql.NullString
}

func main() {
	// Koneksi database
	dbHost := "localhost"
	dbPort := "5432"
	dbName := "gold_pocket"
	dbUser := "postgres"
	dbPassword := "postgres"
	dataSourceName :=
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser,
			dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Connected")
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// Insert
	// _, err = db.NamedExec("INSERT INTO customers (id, first_name,last_name, birth_date, address, status, username, password, email)VALUES (:id, :first_name, :last_name, :birth_date, :address, :status, :username, :password, :email)", &Customers{
	// 	Id:        "2",
	// 	Firstname: "Irfan",
	// 	Lastname:  "Maulana",
	// 	BirtDate:  time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC),
	// 	Address:   "Jalan Jendral Sudirman",
	// 	Status:    1,
	// 	Username:  "irfan",
	// 	Password:  "irfan123",
	// 	Email:     "zKqF@example.com",
	// })

	// if err != nil {
	// 	log.Fatalln(err)
	// } else {
	// 	log.Print("Successfully insert data")
	// }

	// cara kedua
	layoutFormat := "2006-01-02"
	birthDate := "1999-02-22"
	birthDateValue, _ := time.Parse(layoutFormat, birthDate)
	newCustomer := map[string]interface{}{
		"id":         "A009",
		"first_name": "Sun",
		"last_name":  "Mina",
		"birth_date": birthDateValue,
		"address":    "Kore",
		"status":     1,
	}
	_, err = db.NamedExec(`INSERT INTO customers (id, first_name, last_name, birth_date, address, status) VALUES(:id, :first_name, :last_name, :birth_date, :address, :status)`, newCustomer)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Print("Successfully insert data")
	}

	// Select
	customers := []Customers{}
	db.Select(&customers, "SELECT * FROM customers ORDER BY id ASC")
	fmt.Println(customers)
}
