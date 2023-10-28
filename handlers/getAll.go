package handlers

import (
	"goApi/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	dbArr, err := db.GetAllNames()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(200, dbArr)
}
