package main

import (
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id       int
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint hit")
	fmt.Fprint(w, "Welcome to homepage")
}

func returnAllProducts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("product endpoint hit")
}

func handleReq() {
	http.HandleFunc("/products", returnAllProducts)
	http.HandleFunc("/", homePage)
	http.ListenAndServe("10000", nil)
}

func main() {
	Products = []Product{
		Product{Id: 1, Name: "chair", Quantity: 100, Price: 100.00},
		Product{Id: 2, Name: "desk", Quantity: 200, Price: 200.00},
	}
	handleReq()
}
