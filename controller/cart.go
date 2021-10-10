package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pysga1996/spooky-cart-service/error"
	"github.com/pysga1996/spooky-cart-service/model"
	"github.com/pysga1996/spooky-cart-service/util"
	"net/http"
	"strconv"
)

func GetCart(c *gin.Context) {
	if id, err := strconv.ParseUint(c.Param("id"), 10, 64); err == nil {
		var createTime sql.NullTime
		var updateTime sql.NullTime
		var status sql.NullByte
		row := DB.QueryRow("SELECT * FROM cart WHERE id = $1", id)
		err = row.Scan(&id, &createTime, &updateTime, &status)
		if err != nil {
			error.NotFound(c, err)
			return
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
	} else {
		error.BadRequest(c, err)
	}
}
