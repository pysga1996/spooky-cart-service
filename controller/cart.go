package controller

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/pysga1996/spooky-cart-service/constant"
	"github.com/pysga1996/spooky-cart-service/middleware"
	"github.com/pysga1996/spooky-cart-service/model"
	"github.com/pysga1996/spooky-cart-service/util"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetCart(c *gin.Context) {
	WithOutTransaction(c, getCart)
}

func AddCartProduct(c *gin.Context) {
	WithTransaction(c, addCartProduct)
}

func UpdateCartProduct(c *gin.Context) {
	WithTransaction(c, updateCartProduct)
}

func DeleteCartProduct(c *gin.Context) {
	WithTransaction(c, deleteCartProduct)
}

func getCart(c *gin.Context) {
	currentUser := c.GetString(constant.UID)
	if len(currentUser) == 0 {
		middleware.Unauthorized(c, errors.New("not logged in yet"))
		return
	}
	var err error
	var id uint64
	var createTime sql.NullTime
	var updateTime sql.NullTime
	var status sql.NullByte
	var stmt *sql.Stmt
	stmt, err = DB.Prepare("SELECT id, create_time, update_time, status FROM cart WHERE username = $1 AND status = $2")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	row := stmt.QueryRow(currentUser, constant.STATUS_ACTIVE)
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

	cart := new(model.Cart)
	cartProduct := GetCartProducts(c, id)
	cart.SetId(&id)
	cart.SetProductQuantity(cartProduct)
	cart.SetCreateTime(util.GetNullableTime(&createTime))
	cart.SetUpdateTime(util.GetNullableTime(&updateTime))
	cart.SetStatus(util.GetNullableByte(&status))
	c.JSON(http.StatusOK, cart)
}

func CreateCart(c *gin.Context, username string) (success bool) {
	var tx *sql.Tx
	var err error
	if tx, err = DB.Begin(); err != nil {
		middleware.InternalServer(c, err)
		return false
	}
	if err = SetSchema(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
	result, err := tx.Exec("INSERT INTO cart(create_time, status, username) VALUES ($1, $2, $3)", time.Now(), constant.STATUS_ACTIVE, username)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			middleware.InternalServer(c, err)
			return false
		}
		middleware.InternalServer(c, err)
		return false
	}
	var count int64
	if count, err = result.RowsAffected(); err != nil {
		middleware.InternalServer(c, err)
		return false
	}
	log.Printf("Rows affected = %d\n", count)
	if err = tx.Commit(); err != nil {
		middleware.InternalServer(c, err)
		return false
	}
	return true
}

func GetCartProducts(c *gin.Context, cartId uint64) (cartProducts *map[string]*uint8) {
	stmt, err := DB.Prepare("SELECT cart_id, product_code, quantity FROM cart_product WHERE cart_id = $1 AND status = $2")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	rows, err := stmt.Query(cartId, constant.STATUS_ACTIVE)
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	m := make(map[string]*uint8)
	for rows.Next() {
		var id uint64
		var productCode sql.NullString
		var quantity sql.NullByte
		err = rows.Scan(&id, &productCode, &quantity)
		if err != nil {
			middleware.InternalServer(c, err)
			return
		}
		m[*util.GetNullableString(&productCode)] = util.GetNullableByte(&quantity)
	}
	return &m
}

func addCartProduct(c *gin.Context) {
	currentUser := c.GetString(constant.UID)
	if len(currentUser) == 0 {
		middleware.Unauthorized(c, errors.New("not logged in yet"))
		return
	}
	var err error
	var id uint64
	var stmt *sql.Stmt
	stmt, err = DB.Prepare("SELECT id FROM cart WHERE username = $1 AND status = $2")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	row := stmt.QueryRow(currentUser, constant.STATUS_ACTIVE)
	err = row.Scan(&id)
	var newCartProducts []*model.CartProduct
	// Call BindJSON to bind the received JSON to
	if err = c.BindJSON(&newCartProducts); err != nil {
		middleware.BadRequest(c, err)
		return
	}
	stmt, err = DB.Prepare(`INSERT INTO cart_product (cart_id, product_code, quantity) VALUES ($1, $2, $3) ON CONFLICT ON CONSTRAINT cart_product_pk DO UPDATE SET quantity = $3, status = $4`)
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	var result sql.Result
	var count int64 = 0
	for _, product := range newCartProducts {
		result, err = stmt.Exec(id, *product.ProductCode, *product.Quantity, constant.STATUS_ACTIVE)
		var insertCount int64
		if insertCount, err = result.RowsAffected(); err != nil {
			middleware.InternalServer(c, err)
			return
		}
		product.CartId = &id
		count += insertCount
	}
	log.Printf("Rows inserted count: %d", count)
	result, err = DB.Exec("UPDATE cart SET update_time = $2 WHERE id = $1", id, time.Now())
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newCartProducts)
}

func updateCartProduct(c *gin.Context) {
	currentUser := c.GetString(constant.UID)
	if len(currentUser) == 0 {
		middleware.Unauthorized(c, errors.New("not logged in yet"))
		return
	}
	var err error
	var id uint64
	var stmt *sql.Stmt
	stmt, err = DB.Prepare("SELECT id FROM cart WHERE username = $1 AND status = $2")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	row := stmt.QueryRow(currentUser, constant.STATUS_ACTIVE)
	err = row.Scan(&id)
	var newCartProducts []*model.CartProduct
	// Call BindJSON to bind the received JSON to
	if err = c.BindJSON(&newCartProducts); err != nil {
		middleware.BadRequest(c, err)
		return
	}
	stmt, err = DB.Prepare("UPDATE cart_product SET quantity = $3 WHERE cart_id = $1 AND product_code = $2 AND status = $4")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	var result sql.Result
	var count int64 = 0
	for _, product := range newCartProducts {
		result, err = stmt.Exec(id, *product.ProductCode, *product.Quantity, constant.STATUS_ACTIVE)
		var insertCount int64
		if insertCount, err = result.RowsAffected(); err != nil {
			middleware.InternalServer(c, err)
			return
		}
		product.CartId = &id
		count += insertCount
	}
	log.Printf("Rows updated count: %d", count)
	result, err = DB.Exec("UPDATE cart SET update_time = $2 WHERE id = $1", id, time.Now())
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	c.IndentedJSON(http.StatusCreated, newCartProducts)
}

func deleteCartProduct(c *gin.Context) {
	currentUser := c.GetString(constant.UID)
	if len(currentUser) == 0 {
		middleware.Unauthorized(c, errors.New("not logged in yet"))
		return
	}
	var err error
	var id uint64
	var stmt *sql.Stmt
	stmt, err = DB.Prepare("SELECT id FROM cart WHERE username = $1 AND status = $2")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	row := stmt.QueryRow(currentUser, constant.STATUS_ACTIVE)
	err = row.Scan(&id)
	productCodesStr := strings.TrimSpace(c.Query("product-code"))
	productCodes := strings.Split(productCodesStr, ",")
	if len(productCodes) == 0 || len(productCodes[0]) == 0 {
		middleware.BadRequest(c, errors.New("missing parameter product-code"))
		return
	}
	stmt, err = DB.Prepare("UPDATE cart_product SET status = $3 WHERE cart_id = $1 AND product_code = ANY($2)")
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	var result sql.Result
	var count int64
	result, err = stmt.Exec(id, pq.Array(productCodes), constant.STATUS_REMOVED)
	if count, err = result.RowsAffected(); err != nil {
		middleware.InternalServer(c, err)
		return
	}
	log.Printf("Rows updated count: %d", count)
	result, err = DB.Exec("UPDATE cart SET update_time = $2 WHERE id = $1", id, time.Now())
	if err != nil {
		middleware.InternalServer(c, err)
		return
	}
	c.Status(http.StatusOK)
}
