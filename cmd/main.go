package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":3000",
		Handler: rotas(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func rotas() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/createProduct", CreateProduct)
	mux.HandleFunc("/showProduct", ShowProduct)
	mux.HandleFunc("/updateProduct", UpdateProduct)
	mux.HandleFunc("/deleteProduct", DeleteProduct)

	return mux
}
