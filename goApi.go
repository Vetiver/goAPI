
package main

import "github.com/gin-gonic/gin"

import "goApi/handlers"



func main() {
	r := gin.Default()
	v1 := r.Group("/hello")
	{
		v1.GET("/any", handlers.Any)
		v1.GET("/vi", handlers.Vi)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}