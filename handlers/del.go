package handlers

import (
	"github.com/gin-gonic/gin"
	"goApi/db"
)

func Del(c *gin.Context) {
	c.JSON(200, db.DeliteById())
}