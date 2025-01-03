package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/klimenkokayot/calc-api-go/pkg/calculation"
)

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

//////////////////////////////////////////////

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, NewJsonResponse(ErrMethodNotAllowed))
		return
	}

	request := Request{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, NewJsonResponse(ErrBadRequest))
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
