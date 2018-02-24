package common

import (
	"fmt"
	"time"
)

// Program holds information about bonus point programs
type Program struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Companies   []string  `json:"companies,omitempty"`
	Token       string    `json:"token,omitempty"`
	Shop        string    `json:"shop,omitempty"`
	Description string    `json:"description,omitempty"`
	Address     string    `json:"address,omitempty"`
	StartDate   time.Time `json:"startdate,omitempty"`
	EndDate     time.Time `json:"enddate,omitempty"`
	Image       string    `json:"image,omitempty"`
}

// Print prints the program
func (p *Program) Print() {
	fmt.Println("ID: " + p.ID)
	fmt.Println("Name: " + p.Name)

	for _, c := range p.Companies {
		fmt.Println(c)
	}

	fmt.Println("Token: " + p.Token)
	fmt.Println("Shop: " + p.Shop)
	fmt.Println("Description: " + p.Description)
	fmt.Println("Address: " + p.Address)
	fmt.Println("StartDate: " + p.StartDate.String())
	fmt.Println("EndDate: " + p.EndDate.String())
}

// StellarAddress holds a public/private key combo
type StellarAddress struct {
	Seed    string
	Address string
}
