package middleware

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

func RequestLogger(next http.Handler, filename string) http.Handler {
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("RequestLogger cant create file with name: %s", filename)
		return next
	}
	log.SetOutput(file)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := io.ReadAll(r.Body)
		rdr := io.NopCloser(bytes.NewBuffer(buf))

		request := string(buf)

		r.Body = rdr

		log.Println(request, log.Ldate)

		defer func() {
			if err := recover(); err != nil {
				log.Println("recovered RequestLogger", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})

}
