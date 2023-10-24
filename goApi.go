package main

import "github.com/gin-gonic/gin"

import "goApi/handlers"

func main() {
	r := gin.Default()
	v1 := r.Group("/name")
	{
		v1.GET("/insert", handlers.InsertName)
		v1.GET("/getNames", handlers.GetAll)
		v1.GET("/delName", handlers.Del)
		v1.GET("/update", handlers.UpdateNameById)
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
