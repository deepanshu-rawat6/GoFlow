package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	// Creating a new server mux for handling routes
	mux := http.NewServeMux()

	// A simple route to return "Hello World" on path "/"
	mux.HandleFunc("/", s.HelloWorldHandler)

	// Sending repsonse of "Hello from Botgauge" on path "/hello"
	mux.HandleFunc("/hello", s.HelloFromBotgaugeHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// Handler function to print "Hello from Botgauge", with JSON formatting
func (s *Server) HelloFromBotgaugeHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello from Botgauge"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
