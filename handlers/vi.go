package handlers

import (
	"github.com/gin-gonic/gin"
	"goApi/db"
)

func Vi(c *gin.Context) {
	c.JSON(200, db.GetAllNames())
}
