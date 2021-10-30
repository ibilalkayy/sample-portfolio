package database

// Importing the libraries
import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ibilalkayy/Portfolio/middleware"
)

// Building connection with the database
// Using db_user, db_password, db_address and db_db to send data into it
// Establishing the connection and showing the success message
func Connection(name, email, subject, body string) {
	db_user := middleware.LoadEnvVariable("DB_USER")
	db_password := middleware.LoadEnvVariable("DB_PASSWORD")
	db_address := middleware.LoadEnvVariable("DB_ADDRESS")
	db_db := middleware.LoadEnvVariable("DB_DB")
	s := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_password, db_address, db_db)
	db, err := sql.Open("mysql", s)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("Database is successfully connected")

	CreateTable(db)
	InsertData(db, name, email, subject, body)
}

// Creating a table by calling the person.SQL file
// Splitting the string and calling the first query
// Taking the query and executing it
func CreateTable(db *sql.DB) {
	query, err := ioutil.ReadFile("database/person.SQL")
	if err != nil {
		log.Fatal(err)
	}

	requests := strings.Split(string(query), ";")[0]

	stmt, err := db.Prepare(requests)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// Inserting the data into the database
// Executing the file and taking the name, email, subject, and body as values
func InsertData(db *sql.DB, name, email, subject, body string) {
	q := "INSERT INTO Person(names, emails, subjects, bodies) VALUES(?, ?, ?, ?)"
	insert, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	defer insert.Close()

	_, err = insert.Exec(name, email, subject, body)
	if err != nil {
		log.Fatal(err)
	}
}
