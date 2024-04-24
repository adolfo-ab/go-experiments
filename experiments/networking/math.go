package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}

type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	// Decode and validate
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}

	if err := dec.Decode(req); err != nil {
		log.Printf("error: bad JSON: %s", err)
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	if !strings.Contains("+-*/", req.Op) {
		log.Printf("error: bad operator: %q", req.Op)
		http.Error(w, "unknown operator", http.StatusBadRequest)
		return
	}

	// Execute operation
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "division by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("unknown operator: %s", req.Op)
	}

	// Encode results
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		// Can't return error to client here
		log.Printf("can't encode %v - %s", resp, err)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/math", mathHandler)

	addr := ":8080"
	log.Printf("server ready on %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
