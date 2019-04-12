package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/chargen"
	"github.com/ironarachne/random"
)

func getCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var newChar chargen.Character

	random.SeedFromString(id)

	newChar = chargen.GenerateCharacter()

	json.NewEncoder(w).Encode(newChar)
}

func getCharacterRandom(w http.ResponseWriter, r *http.Request) {
	var newChar chargen.Character

	rand.Seed(time.Now().UnixNano())

	newChar = chargen.GenerateCharacter()

	json.NewEncoder(w).Encode(newChar)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", getCharacterRandom)
	r.Get("/{id}", getCharacter)

	fmt.Println("Character Generator API is online.")
	log.Fatal(http.ListenAndServe(":9798", r))
}
