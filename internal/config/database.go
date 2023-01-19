package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/thanh-vt/splash-inventory-service/internal/controller"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strconv"
)

func ConnectDatabase() error {
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
	db, err := gorm.Open(postgres.Open(psqlConn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	controller.DB = db
	if err != nil {
		return err
	}

	fmt.Println("Connected!")
	return err
}
