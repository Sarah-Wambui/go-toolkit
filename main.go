package main
import (
	"encoding/json"
	"net/http"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
        {ID: 1, Name: "Laptop", Price: 1200.50},
        {ID: 2, Name: "Headphones", Price: 150.99},
        {ID: 3, Name: "Keyboard", Price: 89.90},
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/products", productHandler)
	http.ListenAndServe(":8080", nil)
}