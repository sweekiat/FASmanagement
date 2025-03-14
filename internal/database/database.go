package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func ConnectDB() {
        envErr := godotenv.Load()
        if envErr != nil {
                log.Fatal("Error loading .env file")
        }
        fmt.Println("Connecting to the database")
        cfg := mysql.Config{
        User:   os.Getenv("RDS_USER"),
        Passwd: os.Getenv("RDS_PASSWORD"),
        Net:    "tcp",
        Addr:   os.Getenv("RDS_HOST") + ":" + os.Getenv("RDS_PORT"),
        DBName: os.Getenv("RDS_DB_NAME"),
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

func GetDB() *sql.DB {
        return db
}

func CloseDB() {
        db.Close()
}