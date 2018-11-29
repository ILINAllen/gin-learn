package main

import (
	"myapis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", myapis.IndexApi)

	router.POST("/person", myapis.AddPersonApi)

	router.GET("/persons", myapis.GetPersonsApi)

	router.GET("/person/:id", myapis.GetPersonApi)

	router.PUT("/person/:id", myapis.ModPersonApi)

	router.DELETE("/person/:id", myapis.DelPersonApi)

	return router
}
