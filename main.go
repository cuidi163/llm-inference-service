package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 定义输入输出结构体
type RequestBody struct {
	Prompt string `json:"prompt"`
}

type ResponseBody struct {
	Response string `json:"response"`
}

// 模拟推理过程的函数
func inference(prompt string) string {
	// 这里应该调用vLLM的相关API进行实际推理
	return "This is a response from the model."
}

func handler(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := ResponseBody{Response: inference(requestBody.Prompt)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/infer", handler)

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
