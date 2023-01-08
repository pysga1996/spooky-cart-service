package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/thanh-vt/spooky-cart-service/internal/middleware"
)

func SetSchema() error {
	_, err := DB.Exec("SET search_path TO spooky_cart, public")
	return err
}

func WithTransaction(c *gin.Context, handler gin.HandlerFunc) {
	var err error
	var tx *sql.Tx
	if tx, err = DB.Begin(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
	if err = SetSchema(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
	handler(c)
	if err = tx.Commit(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
}

func WithOutTransaction(c *gin.Context, handler gin.HandlerFunc) {
	var err error
	if err = SetSchema(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
	handler(c)
}
