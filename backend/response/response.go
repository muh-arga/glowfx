package response

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"success": true,
		"message": "success",
		"data":    data,
	})
}

func Error(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"success": false,
		"message": message,
	})
}

func ValidationError(c *gin.Context, errs interface{}) {
	c.JSON(422, gin.H{
		"success": false,
		"message": "validation failed",
		"errors":  errs,
	})
}
