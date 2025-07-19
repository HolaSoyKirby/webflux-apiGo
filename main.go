package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Id          int     `json:"productId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Customer struct {
	Id     int    `json:"customerId"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

var products = map[int]Product{
	1: {Id: 1, Name: "Laptop", Description: "Laptop bien chida", Price: 999.99},
	2: {Id: 2, Name: "Mouse", Description: "Mouse bien mouse", Price: 25.50},
	3: {Id: 3, Name: "Teclado", Description: "Teclado mecanico", Price: 70.00},
}

var customers = map[int]Customer{
	1: {Id: 1, Name: "Juan Pérez", Active: true},
	2: {Id: 2, Name: "María López", Active: false},
}

func productListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "No se pudo leer el body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var ids []int
	if err := json.Unmarshal(body, &ids); err != nil {
		http.Error(w, "JSON inválido: se esperaba un array de enteros", http.StatusBadRequest)
		return
	}

	var result []Product
	for _, id := range ids {
		if p, ok := products[id]; ok {
			result = append(result, p)
		}
	}

	jsonResponse(w, result, http.StatusOK)
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/customer/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID de cliente inválido", http.StatusBadRequest)
		return
	}

	customer, found := customers[id]
	if !found {
		http.Error(w, "Cliente no encontrado", http.StatusNotFound)
		return
	}

	jsonResponse(w, customer, http.StatusOK)
}

func jsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/product/get-list-of-products", productListHandler)
	mux.HandleFunc("/api/customer/", customerHandler)

	log.Println("✅ API corriendo en http://localhost:8082")
	if err := http.ListenAndServe(":8082", mux); err != nil {
		log.Fatal(err)
	}
}
