package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goApi/db"
)


func (h BaseHandler) UpdateNameById(c *gin.Context) {
	var user db.Data
	err := c.BindJSON(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	user, err = h.db.UpdateName(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusOK, user)
}
