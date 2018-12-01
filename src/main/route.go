package main

import (
	"myApis"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", myApis.IndexApi)

	router.POST("/person", myApis.AddPersonApi)

	router.GET("/persons", myApis.GetPersonsApi)

	router.GET("/person/:id", myApis.GetPersonApi)

	router.PUT("/person/:id", myApis.ModPersonApi)

	router.DELETE("/person/:id", myApis.DelPersonApi)

	return router
}
