package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"github.com/gorilla/mux"
)
func readFile(w http.ResponseWriter, r *http.Request, filePath string){
	w.Header().Set("Content-Type", "application/json")
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "Error opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileData := make([]byte, fileInfo.Size())
	_, err = file.Read(fileData)
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(fileData)
}
func handlePost1kbText200ms(w http.ResponseWriter, r *http.Request) {
	delayMilliseconds := 200
	time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
	readFile(w,r,"1Kb.txt")
}
func handlePost1kbJSON200ms(w http.ResponseWriter, r *http.Request) {
	delayMilliseconds := 200
	time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
	readFile(w,r,"1Kb.json")
}


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/cgi-bin/text/1kb_text_200ms", handlePost1kbText200ms).Methods("POST")
	r.HandleFunc("/cgi-bin/json/1kb_text_200ms", handlePost1kbJSON200ms).Methods("POST")


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
