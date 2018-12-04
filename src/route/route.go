package route

import (
	"myApis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", myApis.IndexApi)

	router.GET("/person/add", myApis.ShowHtmlPage)

	router.POST("/person/add", myApis.AddPersonApi)

	router.GET("/person/list", myApis.GetPersonsApi)

	router.GET("/person/get/:id", myApis.GetPersonApi)

	router.PUT("/person/update/:id", myApis.ModPersonApi)

	router.DELETE("/person/delete/:id", myApis.DelPersonApi)

	//渲染html页面
	router.LoadHTMLGlob("views/*.html")

	return router
}
