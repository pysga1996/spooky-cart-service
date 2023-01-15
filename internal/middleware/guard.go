package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/thanh-vt/splash-inventory-service/internal/constant"
)

func HandleGuard(c *gin.Context) {
	currentUser := c.GetString(constant.UID)
	if len(currentUser) == 0 {
		Unauthorized(c, errors.New("not logged in yet"))
		return
	}
}
