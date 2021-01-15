//Product Storage Center
package main

import (
	"encoding/json"
	"math/rand"

	//route handlers
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Defining the model(structure)
//product id , product type and features
type Product struct {
	ID       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	Features string `json:"features,omitempty"`
}

var products []Product

//to get all the products
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

//add a new product
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(product)
	product.ID = strconv.Itoa(rand.Intn(1000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(product)
}

//to get a specific product using id
func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

//to update a product using id
func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(product)
			product.ID = params["id"]
			products = append(products, product)
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	json.NewEncoder(w).Encode(products)
}

//to delete a product using id
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(products)
}
func main() {
	//Initializing the router
	router := mux.NewRouter()
	products = append(products, Product{ID: "441", Type: "SmartPhone", Features: "Samsung Galaxy S20+ with 8Gb ram and 256Gb internal storage."})
	products = append(products, Product{ID: "442", Type: "Smartwatch", Features: "Apple smartwatch with heart rate monitor and gps."})
	products = append(products, Product{ID: "443", Type: "Laptop", Features: "Lenonvo ideapad with 8gb ram and 1tb Hdd and intelcorei5 processor"})
	//Creating endpoints
	router.HandleFunc("/products", getProducts).Methods("GET")
	router.HandleFunc("/products", createProduct).Methods("POST")
	router.HandleFunc("/products/{id}", getProduct).Methods("GET")
	router.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
