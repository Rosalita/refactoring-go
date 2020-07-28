package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/text/currency"
)

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func format(amount float64) string {
	return fmt.Sprintf("%+v", currency.USD.Amount(amount))
}

func main() {

	playsFile, err := ioutil.ReadFile("plays.json")
	if err != nil {
		fmt.Println(err)
	}

	var plays map[string]Play
	if err := json.Unmarshal(playsFile, &plays); err != nil {
		fmt.Println(err)
	}

	invoiceFile, err := ioutil.ReadFile("invoices.json")
	if err != nil {
		fmt.Println(err)
	}

	var invoice Invoice
	if err := json.Unmarshal(invoiceFile, &invoice); err != nil {
		fmt.Println(err)
	}

	statement(invoice, plays)

}

func statement(invoice Invoice, plays map[string]Play) {
	// totalAmount := 0
	// volumeCredits := 0
	// result := fmt.Sprintf("Statement for %s \n", invoice.Customer)

	fmt.Println(invoice.Performances)
}
