package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{"error": err.Error()})
}
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{"success": true, "data": data, "error": nil})
}
