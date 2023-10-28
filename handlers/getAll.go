package handlers

import (
	"goApi/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type UserGet struct {
	Parce []db.Data `json:"parce"`
 }

func GetAll(c *gin.Context) {
	dbArr, err := db.GetAllNames()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var parce UserGet
	parce.Parce = dbArr

	c.JSON(http.StatusOK, parce)
}
