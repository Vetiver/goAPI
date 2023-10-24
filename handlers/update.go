package handlers

import (
	"github.com/gin-gonic/gin"
	"goApi/db"
)

func UpdateNameById(c *gin.Context) {
	c.JSON(200, db.UpdateName())
}
