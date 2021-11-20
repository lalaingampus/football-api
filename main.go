package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Player represents player model
type Player struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Position string `json:"position"`
}

// Player represents player model
type Team struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Players     *Player `json:"players"`
}

var teams []Team
var players []Player

// Get all returns list of Teams
func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

// Delete Teams by id
func deleteTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range teams {
		if item.ID == params["id"] {
			teams = append(teams[:index], teams[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(teams)
}

// Post creates a new team with nest struct player
func createTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var team Team
	_ = json.NewDecoder(r.Body).Decode(&team)
	team.ID = strconv.Itoa(rand.Intn(10000000))
	teams = append(teams, team)
	json.NewEncoder(w).Encode(team)
}

// GetByID returns a team by id
func getTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range teams {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// GetByID returns a player by id
func getPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range players {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Update value of team by ID
func updateTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range teams {
		if item.ID == params["id"] {
			teams = append(teams[:index], teams[index+1:]...)
			var team Team
			_ = json.NewDecoder(r.Body).Decode(&team)
			team.ID = params["id"]
			teams = append(teams, team)
			json.NewEncoder(w).Encode(team)
			return
		}
	}
}

func main() {

	r := mux.NewRouter()

	teams = append(teams, Team{ID: "1", Name: "leo", Description: "wow", Location: "wlw", Players: &Player{ID: "1", Name: "kol", Nickname: "loe", Position: "lop"}})

	// HandleFunc initialize endpoints Teams
	r.HandleFunc("/team", getTeams).Methods("GET")
	r.HandleFunc("/team/{id}", getTeam).Methods("GET")
	r.HandleFunc("/player/{id}", getPlayer).Methods("GET")
	r.HandleFunc("/team", createTeam).Methods("POST")
	r.HandleFunc("/team/{id}", updateTeam).Methods("PUT")
	r.HandleFunc("/team/{id}", deleteTeam).Methods("DELETE")

	// For print server running
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
