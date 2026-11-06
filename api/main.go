package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"status": "ok",
			"msg":    "API rodando no container",
			"hora":	  time.Now(),
		})
	})

	fmt.Println("API na porta 3000")
	http.ListenAndServe(":3000", nil)
}