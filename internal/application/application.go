package application

import (
	"fmt"
	"net/http"
	"os"

	"github.com/klimenkokayot/calc-api-go/http/server"
)

type Config struct {
	Port string
}

type Application struct {
	Config *Config
}

func ConfigFromEnv() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{
		Port: port,
	}
}

func NewApplication() *Application {
	return &Application{
		Config: ConfigFromEnv(),
	}
}

func (app *Application) StartServer() error {
	handler := server.Run()
	// Получаем серверок и настраиваем через конфиг и все чики пуки
	fmt.Println("server started at port ", app.Config.Port)
	http.ListenAndServe(":"+app.Config.Port, handler)
	return nil
}
