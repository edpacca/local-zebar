package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func handleError(w http.ResponseWriter, message string, code int) {
	http.Error(w, message, code)
	log.Printf("%v (%v)", message, code)
}

func getVLCStatus(w http.ResponseWriter, r *http.Request) {
	vlcEndpoint := os.Getenv("VLC_ENDPOINT")
	vlcPassword := os.Getenv("VLC_PASSWORD")

	if vlcEndpoint == "" || vlcPassword == "" {
		errorMsg := "Missing environment variables"
		handleError(w, errorMsg, http.StatusInternalServerError)
		return
	}

	// Build request to VLC with basic authentication
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("GET", vlcEndpoint, nil)
	if err != nil {
		handleError(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Set Basic Auth (username is empty, VLC only uses a password)
	req.SetBasicAuth("", vlcPassword)

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		handleError(w, "Failed to connect to VLC", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	vlcResponseMessage := fmt.Sprintf("VLC Response - Status: %d\n", resp.StatusCode)
	log.Println(vlcResponseMessage)

	if resp.StatusCode != http.StatusOK {
		handleError(w, vlcResponseMessage, resp.StatusCode)
		return
	}

	var responseData map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		handleError(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	// Forward VLC response
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Origin, Accept")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(responseData)
}

func main() {
	port := "3999"
	if len(os.Args) > 1 {
		port = os.Args[1] // Allow setting port via command line
	}

	http.HandleFunc("/vlc-status", getVLCStatus)

	fmt.Printf("Proxy server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
