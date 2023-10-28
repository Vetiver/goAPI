package handlers

import (
	"goApi/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
    Name string `json:"name"`
}

func InsertName(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, db.Insert(user.Name))
}
