package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

func handlePostRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received POST request")
}
func handlePostRequest2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Received POST request")
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/post", handlePostRequest).Methods("POST")
	r.HandleFunc("/post2", handlePostRequest2).Methods("POST")


	httpServer := &http.Server{
		Addr:         ":8080", // HTTP port
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	httpsServer := &http.Server{
		Addr:         ":8443", // HTTPS port
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		fmt.Println("Starting HTTP server on port 8080...")
		err := httpServer.ListenAndServe()
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()

	go func() {
		fmt.Println("Starting HTTPS server on port 8443...")
		err := httpsServer.ListenAndServeTLS("cert.pem", "key.pem")
		if err != nil {
			fmt.Println("Error starting HTTPS server:", err)
		}
	}()

	select {}
}
