package main

import (
	"log"
	"net/http"
	"os/exec"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/meshi", func(r chi.Router) {
		r.Get("/", getMeshi) // GET /meshi/
	})

	http.ListenAndServe(":8080", r)
}

func getMeshi(w http.ResponseWriter, r *http.Request) {
	res, err := exec.Command("meshi").Output()
	if err != nil {
		log.Fatal(err)
	}
	w.Write(res)
}
