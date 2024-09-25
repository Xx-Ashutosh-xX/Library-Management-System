package database

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    // Change the connection string according to your MySQL configuration
    connStr := "root:12345678@tcp(localhost:3306)/library" 
    DB, err = sql.Open("mysql", connStr)
    if err != nil {
        log.Fatal(err)
    }

    // Check if the connection is established
    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
}
