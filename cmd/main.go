package main

import "github.com/DevAthhh/upmark/internal/handlers"

func main() {
	router := handlers.Handle()
	router.Run(":8000")
}
