package error

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CheckErrorShutdown(err error) {
	if err != nil {
		panic(err)
	}
}

func BadRequest(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Status(http.StatusBadRequest)
	c.Abort()
}

func NotFound(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Status(http.StatusNotFound)
	c.Abort()
}

func InternalServer(c *gin.Context, err error) {
	_ = c.Error(err)
	c.Status(http.StatusInternalServerError)
	c.Abort()
}

func Handle(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(-1, gin.H{
			"message":   c.Errors[0].Error(),
			"timestamp": time.Now(),
			"path":      c.Request.RequestURI,
		}) // -1 == not override the current error code
	}
}
