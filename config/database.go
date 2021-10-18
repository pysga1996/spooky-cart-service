package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/pysga1996/spooky-cart-service/middleware"
	"os"
	"strconv"
)

func ConnectDatabase() (db *sql.DB) {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	sslMode := os.Getenv("DB_SSL_MODE")
	// connection string
	psqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)

	// open database
	db, err := sql.Open("postgres", psqlConn)
	middleware.CheckErrorShutdown(err)

	// check db
	err = db.Ping()
	middleware.CheckErrorShutdown(err)
	_, err2 := db.Exec(`set search_path='spooky_cart'`)
	middleware.CheckErrorShutdown(err2)
	fmt.Println("Connected!")
	return db
}
