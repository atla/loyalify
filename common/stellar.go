package common

import (
	"fmt"
	"log"
	"strconv"

	b "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"
	"github.com/stellar/go/keypair"
)

func generateAccount() (string, string) {
	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
	}

	return pair.Seed(), pair.Address()
}

func listBalance(account string) {
	/*	account, err := horizon.DefaultTestNetClient.LoadAccount(account)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Balances for account:", pair.Address())

		for _, balance := range account.Balances {
			log.Println(balance)
		}*/
}

// ChangeTrust is a replacement for friendbot
func ChangeTrust(source StellarAddress, destination string) {

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: source.Address},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.ChangeTrust(
			b.Destination{AddressOrSeed: destination},
		),
	)

	txe, err := tx.Sign(source.Seed)
	if err != nil {
		panic(err)
	}

	blob, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction (ChangeTrust) posted in ledger:", resp.Ledger)
}

// CreateAccount is a replacement for friendbot
func CreateAccount(id string, fundingAccount StellarAddress) {

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: fundingAccount.Address},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.CreateAccount(
			b.Destination{AddressOrSeed: id},
			b.NativeAmount{Amount: "100.0"},
		),
	)

	txe, err := tx.Sign(fundingAccount.Seed)
	if err != nil {
		panic(err)
	}

	blob, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction (CreateAccount) posted in ledger:", resp.Ledger)
}

func floatToString(input float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input, 'f', 6, 64)
}

// PayLoyaltyPoints pays loyalty points to a user id
func PayLoyaltyPoints(destination string, source StellarAddress, token string, amount float64) {

	tx, err := b.Transaction(
		b.SourceAccount{AddressOrSeed: source.Address},
		b.TestNetwork,
		b.AutoSequence{SequenceProvider: horizon.DefaultTestNetClient},
		b.Payment(
			b.Destination{AddressOrSeed: destination},
			b.CreditAmount{
				Code:   token,
				Issuer: source.Address,
				Amount: floatToString(amount),
			},
		),
	)

	txe, err := tx.Sign(source.Seed)
	if err != nil {
		panic(err)
	}

	blob, err := txe.Base64()

	if err != nil {
		panic(err)
	}

	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction (PayLoyalityPoints) posted in ledger:", resp.Ledger)
}
