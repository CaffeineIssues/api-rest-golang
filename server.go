package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Vovo struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	Idade    int    `json:"idade"`
	Telefone string `json:"telefone"`
}

var vovos []Vovo

func cadastrarVovo(w http.ResponseWriter, r *http.Request) {
	var v Vovo
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	v.ID = len(vovos) + 1
	vovos = append(vovos, v)

	w.WriteHeader(http.StatusCreated)
}

func listarVovos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vovos)
}

func atualizarVovo(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var v Vovo
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	for i, vovo := range vovos {
		if vovo.ID == id {
			v.ID = vovo.ID
			vovos[i] = v
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}

func removerVovo(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.Atoi(parts[2])
	if err != nil {
		http.NotFound(w, r)
		return
	}

	for i, vovo := range vovos {
		if vovo.ID == id {
			vovos = append(vovos[:i], vovos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/cadastrar", cadastrarVovo).Methods("POST")
	http.HandleFunc("/listar", listarVovos).Methods("GET")
	http.HandleFunc("/atualizar/", atualizarVovo).Methods("PUT")
	http.HandleFunc("/remover/", removerVovo).Methods("DELETE")

	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", nil)
}
