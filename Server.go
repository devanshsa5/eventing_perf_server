package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	"github.com/gorilla/mux"
	"strings"
)
func readFile(w http.ResponseWriter, r *http.Request, filePath string){
	
	if strings.Contains(filePath, ".txt") {
		w.Header().Set("Content-Type", "plain/text")
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
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

func handlePost(delayMilliseconds uint16, filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
		readFile(w,r,filePath)
	}
}

func handlePostImage(delayMilliseconds uint16 ) http.HandlerFunc{

return func(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(delayMilliseconds) * time.Millisecond)
	w.Header().Set("Content-Type", "image/jpeg")

	imagePath := "100kb_image.jpg"
	imageFile, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer imageFile.Close()

	imageFileInfo, err := imageFile.Stat()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Allocate a buffer to hold the image data
	imageData := make([]byte, imageFileInfo.Size())

	// Read the image data into the buffer
	_, err = imageFile.Read(imageData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}


	// Write the image data to the response writer
	_, err = w.Write(imageData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
}

func main() {
	r := mux.NewRouter()
	jsonFilePath := "/root/eventing_perf_server/1Kb.json"
	textFilePath := "/root/eventing_perf_server/1Kb.txt"
	


	r.HandleFunc("/cgi-bin/text/1kb_text_200ms", handlePost(200, textFilePath)).Methods("POST")
	r.HandleFunc("/cgi-bin/json/1kb_text_200ms", handlePost(200, jsonFilePath)).Methods("POST")
	r.HandleFunc("/cgi-bin/1kb_text", handlePost(0, textFilePath)).Methods("POST")
	r.HandleFunc("/cgi-bin/json/1kb_text_10s", handlePost(10000, jsonFilePath)).Methods("POST")
	r.HandleFunc("/cgi-bin/image/100kb_image_200ms", handlePostImage(200)).Methods("POST")
	r.HandleFunc("/cgi-bin/json/1kb_text_20ms", handlePost(20, jsonFilePath)).Methods("POST")
	r.HandleFunc("/cgi-bin/json/1kb_text_2s", handlePost(2000, jsonFilePath)).Methods("POST")


	

	httpServer := &http.Server{
		Addr:         ":8080", // HTTP port
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	httpsServer := &http.Server{
		Addr:         ":8443", // HTTPS port
		Handler:      r,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
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
		err := httpsServer.ListenAndServeTLS("/usr/cert/cert.pem", "/usr/cert/key.pem")
		if err != nil {
			fmt.Println("Error starting HTTPS server:", err)
		}
	}()

	select {}
}
