package main

import (
	_ "backendAleshinaFeklistova/api"
	"backendAleshinaFeklistova/internal/app"
	"log"
)

func main() {
	a := app.NewApp("localhost", "8080")
	err := a.StartApp()
	if err != nil {
		log.Fatal("failed to start the application")
	}
}
