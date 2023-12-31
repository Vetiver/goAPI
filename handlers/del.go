package handlers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)





func (h BaseHandler)  Del(c *gin.Context) {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	err = h.db.DeleteById(intId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
