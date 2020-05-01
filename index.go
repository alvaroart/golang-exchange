package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {

	valores, ok := r.URL.Query()["valor"]

	if !ok || len(valores[0]) < 1 {
		fmt.Fprintf(w,"Cotação não informada, passe os valores!")
		return
	}
	valor, err := strconv.ParseFloat(valores[0], 32);
	if err != nil  {
		fmt.Fprintf(w, "Cotação de informar números!")
		return
	}

	cotacao := valor / 5.44
	fmt.Fprintf(w, "Sua cotação: %f!", cotacao)

}

func main() {
	http.HandleFunc("/cotacao", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
