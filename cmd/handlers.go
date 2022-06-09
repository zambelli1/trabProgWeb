package main

import (
	"encoding/json"
	"net/http"
	"trabalho_poo/cmd/mysql"
)

type Product struct {
	Id         int     `json:"id"`
	Nome       string  `json:"nome"`
	Marca      string  `json:"marca"`
	Quantidade int     `json:"quantidade"`
	Preco      float64 `json:"preco"`
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("CreateProduct só aceita requesição do tipo POST"))
		return
	}

	bodyDecoder := json.NewDecoder(r.Body)

	var produto Product
	bodyDecoder.Decode(&produto)

	if produto.Nome == "" || produto.Marca == "" || produto.Quantidade <= 0 || produto.Preco <= 0 {
		w.Write([]byte("Precisa informar Nome, Marca, Quantidade e Preco do produto"))
		return
	}

	result, err := mysql.InsertProduct(produto.Nome, produto.Marca, produto.Quantidade, produto.Preco)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if result {
		w.Write([]byte("Produto criado com sucesso"))
	} else {
		w.Write([]byte("Falha na criação do produto"))
	}
}

func ShowProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("ShowProduct só aceita requesição do tipo GET"))
		return
	}

	var produtos []Product
	result, err := mysql.SelectProducts()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	for result.Next() {
		var produto Product
		result.Scan(&produto.Id, &produto.Nome, &produto.Marca, &produto.Quantidade, &produto.Preco)
		produtos = append(produtos, produto)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("UpdateProduct só aceita requesição do tipo POST"))
		return
	}

	bodyDecoder := json.NewDecoder(r.Body)

	var produto Product
	bodyDecoder.Decode(&produto)

	if produto.Id <= 0 || produto.Nome == "" || produto.Marca == "" || produto.Quantidade <= 0 || produto.Preco <= 0 {
		w.Write([]byte("Precisa informar Id, Nome, Marca, Quantidade e Preco do produto para atualizar"))
		return
	}

	result, err := mysql.UpdateProduct(produto.Id, produto.Nome, produto.Marca, produto.Quantidade, produto.Preco)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if result {
		w.Write([]byte("Produto atualizado com sucesso"))
	} else {
		w.Write([]byte("Produto não encontrado ou nenhum campo foi alterado"))
	}
}