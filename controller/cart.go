package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/pysga1996/spooky-cart-service/constant"
	"github.com/pysga1996/spooky-cart-service/middleware"
	"github.com/pysga1996/spooky-cart-service/model"
	"github.com/pysga1996/spooky-cart-service/util"
	"log"
	"net/http"
	"time"
)

func GetCart(c *gin.Context) {
	currentUser := c.GetString(constant.UID)
	if len(currentUser) == 0 {
		middleware.Unauthorized(c, errors.New("not logged in yet"))
		return
	}
	var id uint64
	var createTime sql.NullTime
	var updateTime sql.NullTime
	var status sql.NullByte
	stmt, err := DB.Prepare("SELECT id, create_time, update_time, status FROM cart WHERE username = $1")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	row := stmt.QueryRow(currentUser)
	err = row.Scan(&id, &createTime, &updateTime, &status)
	if err != nil {
		success := CreateCart(c, currentUser)
		if !success {
			middleware.InternalServer(c, err)
			return
		}
		err = row.Scan(&id, &createTime, &updateTime, &status)
		if err != nil {
			middleware.InternalServer(c, err)
			return
		}
	}
	m := make(map[uint64]uint8)
	m[1] = 54
	m[2] = 87
	m[3] = 110
	cart := new(model.Cart)
	cart.SetId(&id)
	cart.SetProductQuantity(&m)
	cart.SetCreateTime(util.GetNullableTime(&createTime))
	cart.SetUpdateTime(util.GetNullableTime(&updateTime))
	cart.SetStatus(util.GetNullableByte(&status))
	c.JSON(http.StatusOK, cart)
}

func CreateCart(c *gin.Context, username string) (success bool) {
	tx, err := DB.Begin()
	if err != nil {
		middleware.InternalServer(c, err)
		return false
	}
	// Add new cart
	result, err := tx.Exec("INSERT INTO cart(create_time, status, username) VALUES ($1, $2, $3)", time.Now(), 1, username)
	if err != nil {
		if err := tx.Rollback(); err != nil {
			middleware.InternalServer(c, err)
			return false
		}
		middleware.InternalServer(c, err)
		return false
	}
	count, err := result.RowsAffected()
	if err != nil {
		middleware.InternalServer(c, err)
		return false
	}
	log.Printf("Rows affected = %d\n", count)
	if err := tx.Commit(); err != nil {
		middleware.InternalServer(c, err)
		return false
	}
	return true
}
