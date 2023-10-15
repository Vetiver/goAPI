package handlers

import "github.com/gin-gonic/gin"
import "goApi/db"

func Any(c *gin.Context) {
	var result = db.DbStart
	c.JSON(200, result)
}