package main

import (
	db "myDatabase"
	"route"
)

func main() {
	defer db.Engine.Close()
	router := route.InitRouter()
	// router.LoadHTMLFiles("views/index.html")
	// router.Static("../static", "./static")
	router.Static("/static", "./static")
	router.Run(":8000")
}
