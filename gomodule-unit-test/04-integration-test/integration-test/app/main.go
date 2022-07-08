package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", HandleAddInts)
	address := ":9000"
	fmt.Printf("Memulai server at %s\n", address)
	fmt.Println()
	fmt.Println("contoh query: /add?a=2&b=2&authtoken=abcdef123")
	fmt.Println("tokennya adalah 'abcdef123'")
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("error ketika start server: %v", err)
	}
}

type TambahResponse struct {
	Result int `json:"result"`
}

func HandleAddInts(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	paramA := params.Get("a")
	paramB := params.Get("b")
	token := params.Get("authtoken")

	if paramA == "" || paramB == "" || token == "" {
		http.Error(w, "tidak ada parameters", http.StatusBadRequest)
		fmt.Println("tidak ada parameters")
		return
	}

	if token != "abcdef123" {
		http.Error(w, "token tidak valid", http.StatusUnauthorized)
		fmt.Println("tidak ada parameters")
		return
	}

	intA, err := strconv.Atoi(paramA)
	if err != nil {
		http.Error(w, "parameter 'a' harus integer", http.StatusBadRequest)
		fmt.Println("parameter 'a' harus integer")
		return
	}

	intB, err := strconv.Atoi(paramB)
	if err != nil {
		http.Error(w, "parameter 'b' harus integer", http.StatusBadRequest)
		fmt.Println("parameter 'b' harus integer")
		return
	}

	response := TambahResponse{
		Result: intA + intB,
	}

	json, err := json.MarshalIndent(&response, "", " ")
	if err != nil {
		http.Error(w, "error while marshalling", http.StatusInternalServerError)
		fmt.Println("parameter 'b' harus integer")
		return
	}

	fmt.Fprint(w, string(json))
}
