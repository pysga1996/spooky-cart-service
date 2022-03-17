package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/thanh-vt/spooky-cart-service/middleware"
	"os"
	"strconv"
)

func ConnectDatabase() (db *sql.DB) {
	// connection string
	var psqlConn string
	if os.Getenv("DATABASE_URL") != "" {
		psqlConn = os.Getenv("DATABASE_URL")
	} else {
		host := os.Getenv("DB_HOST")
		port, _ := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
		user := os.Getenv("DB_USERNAME")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_DATABASE")
		sslMode := os.Getenv("DB_SSL_MODE")
		psqlConn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)
	}

	// open database
	db, err := sql.Open("postgres", psqlConn)
	middleware.CheckErrorShutdown(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)

	// check db
	err = db.Ping()
	middleware.CheckErrorShutdown(err)
	fmt.Println("Connected!")
	return db
}
