package handlers

import (
	"github.com/gin-gonic/gin"
	"goApi/db"
)

func InsertName(c *gin.Context) {
	c.JSON(200, db.Insert())
}
