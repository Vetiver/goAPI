package main

import "github.com/gin-gonic/gin"

import "goApi/handlers"

func main() {
	r := gin.Default()
	v1 := r.Group("/name")
	{
		v1.POST("/insert", handlers.Any)
		v1.GET("/getNames", handlers.Vi)
		v1.POST("/delName", handlers.Del)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
