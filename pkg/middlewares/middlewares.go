package middlewares

import (
	"io"
	"log"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, _ := io.ReadAll(r.Body)
		request := string(data)

		log.Println(request)

		defer func() {
			if err := recover(); err != nil {
				log.Println("recovered RequestLogger", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})

}
