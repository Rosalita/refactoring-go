package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"strings"

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

	result, err := statement(invoice, plays)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func statement(invoice Invoice, plays map[string]Play) (string, error) {
	totalAmount := 0
	var volumeCredits int

	var result strings.Builder

	result.WriteString(fmt.Sprintf("Statement for %s \n", invoice.Customer))

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]
		thisAmount := 0

		switch play.Type {
		case "tragedy":
			thisAmount = 40000
			if perf.Audience > 30 {
				thisAmount += 1000 * (perf.Audience - 30)
			}
		case "comedy":
			thisAmount = 30000
			if perf.Audience > 20 {
				thisAmount += 10000 + 500*(perf.Audience-20)
			}
			thisAmount += 300 * perf.Audience

		default:
			return "", fmt.Errorf("error: unknown performance type %s", play.Type)
		}

		// add volume credits
		volumeCredits += int(math.Max(float64(perf.Audience)-30, 0))
		// add extra credit for every ten comedy attendees

		if "comedy" == play.Type {
			volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
		}

		// print line for this order
		result.WriteString(fmt.Sprintf("%s: %s (%d seats) \n", play.Name, format(float64(thisAmount/100)), perf.Audience))
		totalAmount += thisAmount

	}
	result.WriteString(fmt.Sprintf("Amount owed is %s\n", format(float64(totalAmount)/100)))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", volumeCredits))
	return result.String(), nil
}
