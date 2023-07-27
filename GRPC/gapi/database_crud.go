package gapi

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func DbInit() *sql.DB {
	// Replace the following connection string with your MySQL credentials and database name
	DB, err = sql.Open("mysql", "research:research@tcp(127.0.0.1:3306)/research")
	if err != nil {
		panic(err.Error())
	}
	// defer DB.Close()

	// Test the connection
	err = DB.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to MySQL database!")
	return DB
}

type DBUser struct {
	username  string
	password  string
	email     string
	full_name string
}

func InsertData(db *sql.DB, user DBUser) {

	insertQuery := "INSERT INTO users (username, email,password,full_name) VALUES (?, ?,?,?)"
	_, err := db.Exec(insertQuery, user.username, user.email, user.password, user.full_name)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data inserted successfully!")
}

func LoginData(db *sql.DB, user DBUser) *DBUser {

	insertQuery := "SELECT * FROM users WHERE username = ? AND password = ?"
	row, err := db.Query(insertQuery, user.username, user.password)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data inserted successfully!")
	var username, email string

	row.Scan(&username, &email)
	return &DBUser{username: username, email: email}
}
