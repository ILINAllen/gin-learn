package main

import (
	db "mydatabase"
)

func main() {
	defer db.Engine.Close()
	router := initRouter()
	router.Run(":8000")
}
