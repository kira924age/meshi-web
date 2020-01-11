package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/meshi", func(r chi.Router) {
		r.Get("/", getMeshi) // GET /meshi/
	})

	http.ListenAndServe(":9990", r)
}

func getMeshi(w http.ResponseWriter, r *http.Request) {
	result, err := exec.Command("/opt/local/bin/meshi").Output()
	if err != nil {
		log.Fatal(err)
	}

	resultStr := strings.TrimRight(string(result), "\n")
	jsonBytes := getMeshiByte(resultStr)
	jsonStr := string(jsonBytes)

	w.Write([]byte(fmt.Sprintf(jsonStr)))
}
