package handlers

import (
	"github.com/gin-gonic/gin"
	"goApi/db"
)

func GetAll(c *gin.Context) {
	c.JSON(200, db.GetAllNames())
}
