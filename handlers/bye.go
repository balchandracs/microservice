package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func NewBye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (h *Bye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Bye Bala")

	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Bala %s", d)
}

