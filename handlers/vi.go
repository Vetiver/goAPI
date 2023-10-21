package handlers

import "github.com/gin-gonic/gin"

func Vi(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "здравствуйте Виктор... я опять забыл вашу фамилию)",
	})
}
