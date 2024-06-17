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

		for _, err := range c.Errors {
			log.Printf("Error: %+v", err.Err)
		}
		ginErr := c.Errors[0]
		e, ok := ginErr.Err.(*exception.Exception)
		if !ok {
			e = exception.ServerError
		}

		c.JSON(e.Status, e)
		c.Abort()
	}
}
