package server

import (
	"encoding/json"
	"log"
	"loyalify/common"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var programs []common.Program
var programAddresses []common.StellarAddress

func init() {
	programs = append(programs, common.Program{
		ID:          "1",
		Name:        "Kaufland Monopoly Campaign",
		Company:     "Kaufland",
		Token:       "KauflandM001",
		Shop:        "http://tothemoon.shop",
		Description: "Collect Kaufland BOnus Points and spend them on fantastic new monopoly games assets",
		Address:     "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
		StartDate:   time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
		EndDate:     time.Date(2018, 3, 1, 0, 0, 0, 0, time.Local),
	})

	programs = append(programs, common.Program{
		ID:          "2",
		Name:        "LIDL Special",
		Company:     "LIDL",
		Token:       "LIDLS001",
		Shop:        "http://tothemoon.shop",
		Description: "Collect LIDL Special points",
		Address:     "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
		StartDate:   time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
		EndDate:     time.Date(2018, 3, 1, 0, 0, 0, 0, time.Local),
	})

	//TODO: Remove seed from code
	programAddresses = append(programAddresses, common.StellarAddress{
		Seed:    "",
		Address: "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
	})
}

// GetPrograms returns the list of programs
func GetPrograms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(programs)
}

// GetProgram returns a single program
func GetProgram(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	for _, item := range programs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

// CreateAccount funds a new account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	account := params["id"]

	log.Println("Funding new account", account)
	// fundingAccount
	common.CreateAccount(account, programAddresses[0])

	w.Write([]byte("Created"))
}

// StartServer runs the api server
func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/programs", GetPrograms).Methods("GET")
	router.HandleFunc("/api/programs/{id}", GetProgram).Methods("GET")
	router.HandleFunc("/api/createAccount/{id}", CreateAccount).Methods("GET")
	//router.HandleFunc("/programs/{id}", CreatePerson).Methods("POST")
	//router.HandleFunc("/programs/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))

}
