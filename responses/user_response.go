package responses

import (
	"github.com/gin-gonic/gin"
)

func Error(status int, message string, c *gin.Context) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
	})
}

func Success(status int, data interface{}, c *gin.Context) {
	c.JSON(status, gin.H{
		"status": status,
		"data":   data,
	})
}

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
