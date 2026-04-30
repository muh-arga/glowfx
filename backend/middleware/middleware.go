package middleware

import (
	"errors"
	"glow-fx/response"
	"glow-fx/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors.Last().Err

		if ve, ok := err.(usecase.ValidationErrors); ok {
			response.ValidationError(c, ve)
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, "data not found")
			return
		}

		response.Error(c, err.Error())
	}
}
