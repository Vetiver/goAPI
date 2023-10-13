
package main

import "github.com/gin-gonic/gin"

import "goApi/heandlers"



func main() {
	r := gin.Default()
	v1 := r.Group("/hello")
	{
		v1.GET("/any", heandlers.Any)
		v1.GET("/vi", heandlers.Vi)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}