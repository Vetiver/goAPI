package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h BaseHandler) GetAll(c *gin.Context) {
	dbArr, err := h.db.GetAllNames()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(fmt.Errorf("тут"))
		
		return
	}

	var parce UserGet
	parce.Parce = dbArr
	fmt.Println(parce)

	c.JSON(http.StatusOK, parce)
}
