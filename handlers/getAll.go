package handlers

import (
	"fmt"
	"goApi/db"
	"net/http"

	"github.com/gin-gonic/gin"
)


type UserGet struct {
	Parce []db.Data `json:"parce"`
 }

func GetAll(c *gin.Context) {
	dbArr, err := db.GetAllNames()
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
