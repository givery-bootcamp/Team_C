package middleware

import (
	"log"
	"myapp/internal/exception"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		log.Printf("%+v", c.Errors)
		ginErr := c.Errors[0]
		e, ok := ginErr.Err.(*exception.Exception)
		if !ok {
			e = exception.ServerError
		}

		c.JSON(e.Status, e)
		c.Abort()
	}
}
