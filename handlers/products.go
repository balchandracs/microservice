package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	data "github.com/balchandracs/microservice/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "marshalling error", http.StatusInternalServerError)
	}
	w.Write(d)

}
