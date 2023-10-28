package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goApi/db"
)

type UpdateData struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
   }

func UpdateNameById(c *gin.Context) {
	var user UpdateData
	err := c.BindJSON(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	c.JSON(200, db.UpdateName(user.Name, user.Id))
}
