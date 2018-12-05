package main

import (
	"encoding/json"
	"moecord-api/moedex"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	moedex.Load()

	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/info/{id}", pokemonRoute)

	http.ListenAndServe(":1337", r)
}

func pokemonRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokeid, err := strconv.Atoi(vars["id"])
	pokeid--

	if pokeid < 0 || len(moedex.Moemons) < pokeid {
		w.WriteHeader(500)
		return
	}

	if err != nil {
		w.WriteHeader(500)
		return
	}

	pokeJSON, _ := json.Marshal(moedex.Moemons[pokeid])

	w.Write(pokeJSON)
}
