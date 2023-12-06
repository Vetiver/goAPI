package handlers

import (
	"goApi/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h BaseHandler) InsertName(c *gin.Context) {
	var user *db.Data
	if err := c.BindJSON(&user); 
	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user, err := h.db.Insert(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
