package application

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/klimenkokayot/calc-api-go/internal/apperrors"
	"github.com/klimenkokayot/calc-api-go/pkg/calculation"
	"github.com/klimenkokayot/calc-api-go/pkg/middlewares"
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

type Request struct {
	Expression string `json:"expression"`
}

// omitempty - поле скипается если empty
type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string  `json:"error,omitempty"`
}

func NewJsonResponse(value interface{}) string {
	response := Response{}
	if _, ok := value.(error); ok {
		response.Error = value.(error).Error()
	}
	if _, ok := value.(float64); ok {
		response.Result = value.(float64)
	}
	data, _ := json.Marshal(response)
	return string(data)
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, NewJsonResponse(apperrors.ErrMethodNotAllowed))
		return
	}

	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, NewJsonResponse(apperrors.ErrBadRequest))
		return
	}

	result, err := calculation.Calc(request.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		io.WriteString(w, NewJsonResponse(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, NewJsonResponse(result))
}

func NewApplication() *Application {
	return &Application{
		Config: ConfigFromEnv(),
	}
}

func (app *Application) StartServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", CalcHandler)

	mux2 := middlewares.RequestLogger(mux)

	fmt.Println("server started at port ", app.Config.Port)
	http.ListenAndServe(":"+app.Config.Port, mux2)
	return nil
}
