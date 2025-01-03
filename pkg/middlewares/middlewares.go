package middlewares

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := io.ReadAll(r.Body)
		rdr := io.NopCloser(bytes.NewBuffer(buf))

		request := string(buf)

		r.Body = rdr

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
