package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// APIエンドポイント
func helloHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		http.Error(w, "パラメータ'n'は1以上の整数で指定してください", http.StatusBadRequest)
		return
	}

	result := make([]string, n)
	for i := 0; i < n; i++ {
		result[i] = fmt.Sprintf("item%d", i+1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	http.HandleFunc("/js/practice.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/js/practice.js")
	})

	// APIエンドポイント
	http.HandleFunc("/api/hello", helloHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
