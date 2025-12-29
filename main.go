package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// 定义输入输出结构体
type RequestBody struct {
	Prompt string `json:"prompt"`
}

type ResponseBody struct {
	Response string `json:"response"`
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/infer", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req RequestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("Bad request: %v", err)
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if req.Prompt == "" {
			http.Error(w, "Prompt is required", http.StatusBadRequest)
			return
		}

		log.Printf("Processing prompt: %s", req.Prompt)
		// 模拟推理延迟（关键！）
		time.Sleep(100 * time.Millisecond)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"result": "mock response for: " + req.Prompt,
		})
	})

	http.HandleFunc("/healthz", healthzHandler)

	// 就绪探针（Ready）
	http.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// 存活探针（Alive）
	http.HandleFunc("/live", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
