package heandlers

import "github.com/gin-gonic/gin"

func Any(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
}