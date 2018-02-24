package server

import (
	"encoding/json"
	"fmt"
	"log"
	"loyalify/common"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var programs []common.Program
var programAddresses []common.StellarAddress
var nextProgramID int64

func init() {
	programs = append(programs, common.Program{
		ID:               "1",
		Name:             "Kaufland Monopoly Campaign",
		Companies:        []string{"LIDL"},
		Token:            "KauflandM001",
		Shop:             "http://1876ab32.ngrok.io/shop_web/",
		Description:      "Collect Kaufland Bonus Points and spend them on fantastic new monopoly games assets",
		Address:          "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
		StartDate:        time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
		EndDate:          time.Date(2018, 3, 1, 0, 0, 0, 0, time.Local),
		Image:            "http://lorempixel.com/g/400/200/",
		BonusPointsSpent: 0,
	})

	programs = append(programs, common.Program{
		ID:               "2",
		Name:             "LIDL Special",
		Companies:        []string{"Kaufland"},
		Token:            "LIDLS001",
		Shop:             "http://1876ab32.ngrok.io/shop_web/",
		Description:      "Collect LIDL Special points",
		Address:          "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
		StartDate:        time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
		EndDate:          time.Date(2018, 3, 1, 0, 0, 0, 0, time.Local),
		Image:            "http://lorempixel.com/g/400/200/",
		BonusPointsSpent: 0,
	})

	programs = append(programs, common.Program{
		ID:               "3",
		Name:             "Paykek",
		Companies:        []string{"Kaufland", "REWE", "MediaMarkt", "LIDL", "EDEKA"},
		Token:            "PKEK001",
		Shop:             "http://1876ab32.ngrok.io/shop_web/",
		Description:      "Collect Paykek points",
		Address:          "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
		StartDate:        time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
		EndDate:          time.Date(2022, 3, 1, 0, 0, 0, 0, time.Local),
		Image:            "http://lorempixel.com/g/400/200/",
		BonusPointsSpent: 0,
	})

	//TODO: Remove seed from code
	programAddresses = append(programAddresses, common.StellarAddress{
		Seed:    "",
		Address: "GC7JELXPLI6OLYUKAOAKCTNDWY75WTUUYCKJO4SG4CH3762FICMOVHQK",
	})

	nextProgramID = 3
}

func getNewID() string {

	max := 0.0
	for _, program := range programs {

		current, _ := strconv.Atoi(program.ID)
		max = math.Max(float64(max), float64(current))
	}

	return strconv.Itoa(int(max + 1))
}

func getRandomCompanyName() string {
	companies := []string{"LIDL", "Kaufland", "Edeka", "REWE"}
	max := len(companies)
	index := rand.Int31n(int32(max))
	return companies[index]
}

// CreateProgram asd
func CreateProgram(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating new Program")
	var program common.Program
	_ = json.NewDecoder(r.Body).Decode(&program)
	program.ID = getNewID()
	program.Address = programAddresses[0].Address
	program.Companies = []string{getRandomCompanyName()}
	programs = append(programs, program)
	json.NewEncoder(w).Encode(programs)
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

// PlaceTransaction asda
type PlaceTransaction struct {
	SourceID      string  `json:"source,omitempty"`
	ProgramID     string  `json:"program,omitempty"`
	Amount        float64 `json:"amount,omitempty"`
	Currency      string  `json:"currency,omitempty"`
	TransactionID string  `json:"transaction,omitempty"`
}

// TransactionPlaced asd
func TransactionPlaced(w http.ResponseWriter, r *http.Request) {

	var placeTransaction PlaceTransaction
	_ = json.NewDecoder(r.Body).Decode(&placeTransaction)

	fmt.Println(placeTransaction)

	for idx, program := range programs {
		if program.ID == placeTransaction.ProgramID {

			loyalityPointsEarned := placeTransaction.Amount * 0.2
			formatted := fmt.Sprintf("%f", loyalityPointsEarned)

			//TODO: track transactionid so no two loyality point payouts happen
			common.PayLoyaltyPoints(placeTransaction.SourceID, programAddresses[0], program.Token, loyalityPointsEarned)

			programs[idx].BonusPointsSpent += int64(loyalityPointsEarned)

			answer := "Created loyality points for program " + program.Name + " points earned: " + formatted + " " + program.Token
			// crreate loyality points
			w.Write([]byte(answer))
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
	router.HandleFunc("/api/programs", CreateProgram).Methods("POST")

	router.HandleFunc("/api/createAccount/{id}", CreateAccount).Methods("GET")
	router.HandleFunc("/api/transactionPlaced", TransactionPlaced).Methods("POST")

	router.PathPrefix("/counter").Handler(http.FileServer(http.Dir("dist/")))
	router.PathPrefix("/manage").Handler(http.FileServer(http.Dir("dist/")))

	//router.HandleFunc("/programs/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
