package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "rkb"
	password = "pgsql2020"
	dbname   = "devdb"
)

// model definition
type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func check(err error, msg string) {
	if err != nil {
		log.Fatalf("unable to open database: %v", err)
		return
	}
	log.Println(msg)
}

func getAllUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM users")
	check(err, "query successful")
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		err = rows.Scan(
			&user.ID,
			&user.Age,
			&user.FirstName,
			&user.LastName,
			&user.Email,
		)
		check(err, "row scan complete")
		users = append(users, user)
	}
	return users
}

func main() {
	dbInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfo)
	check(err, "connected to db")
	defer db.Close()

	err = db.Ping()
	check(err, "ping successful")

	users := getAllUsers(db)
	fmt.Println(users)
}
