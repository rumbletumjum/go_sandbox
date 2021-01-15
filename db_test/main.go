package main

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const (
    host = "localhost"
    port = 5432
    user = "rkb"
    password = "pgsql2020"
    dbname = "devdb"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    psqlInfo := fmt.Sprintf(
        "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    check(err)
    defer db.Close()

    err = db.Ping()
    check(err)

    sqlStatement := `
INSERT INTO users (age, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING id`

    var id int
    stmt, err := db.Prepare(sqlStatement)
    check(err)
    err = stmt.QueryRow(36, "rkb@rkbkr.io", "Ron", "Booker").Scan(&id)
    check(err)

    fmt.Println("New record ID is: ", id)
}


