package config

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/thanh-vt/splash-inventory-service/internal/middleware"
	"os"
	"strconv"
)

var DB *sql.DB

func ConnectDatabase() {
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
		schema := os.Getenv("DB_SCHEMA")
		psqlConn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s", host, port, user, password, dbname, sslMode, schema)
	}

	// open database
	DB, err := sql.Open("postgres", psqlConn)
	// close database
	defer func(db *sql.DB) {
		err := db.Close()
		middleware.CheckErrorShutdown(err)
	}(DB)

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)

	// check db
	err = DB.Ping()
	middleware.CheckErrorShutdown(err)
	fmt.Println("Connected!")
}

//func SetSchema() error {
//	_, err := DB.Exec("SET search_path TO splash_inventory, public")
//	return err
//}

func WithTransaction(c *gin.Context, handler gin.HandlerFunc) {
	var err error
	var tx *sql.Tx
	if tx, err = DB.Begin(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
	handler(c)
	if err = tx.Commit(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
}
