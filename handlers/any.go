package handlers

import (
	"github.com/gin-gonic/gin"
	"goApi/db"
)



func Any(c *gin.Context) {
	c.JSON(200, db.Insert())
}

