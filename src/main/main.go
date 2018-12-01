package main

import (
	db "myDatabase"
)

func main() {
	defer db.Engine.Close()
	router := initRouter()
	router.Run(":8000")
}
