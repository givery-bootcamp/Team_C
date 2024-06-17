package middleware

import (
	"errors"
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
		var e *exception.Exception
		unwrappedErr := ginErr.Err
		for {
			if err, ok := unwrappedErr.(*exception.Exception); ok {
				e = err
				break
			}
			unwrappedErr = errors.Unwrap(unwrappedErr)
			if unwrappedErr == nil {
				e = exception.ServerError
				break
			}
		}

		c.JSON(e.Status, e)
		c.Abort()
	}
}
