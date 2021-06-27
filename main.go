package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
  "os"
  "log"

	"github.com/andybalholm/brotli"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
  "github.com/joho/godotenv"
)

type brotliResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func main() {
	fmt.Println("Hello dude.")

	db := InitDB()
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
				next.ServeHTTP(w, r)
				return
			}
			w.Header().Set("Content-Encoding", "br")

			br := brotli.NewWriter(w)
			defer br.Close()

			brr := brotliResponseWriter{Writer: br, ResponseWriter: w}
			next.ServeHTTP(brr, r)
		})
	})

	// r.HandleFunc("/ts", Ts())
	Route(r, db)

  r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./frontend"))))

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"}}).Handler(r)

  err := godotenv.Load()

  if err != nil {
    log.Fatal("Error loading .env file.")
    return
  }

  serverPort := os.Getenv("SERVER_PORT")

	fmt.Println("Serving on port " + serverPort + "!")

	// Generate typescript definitions
	Ts()
	http.ListenAndServe(":" + serverPort, handler)
}
