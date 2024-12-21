package main

import "github.com/klimenkokayot/calc-api-go/internal/application"

func main() {
	app := application.NewApplication()
	app.StartServer()
}
