// Imagine a company of theatrical players who go out to various
// events performing plays. Typically a customer will request a play and the 
// company charges them based on the size of the audience and the kind
// of play they perform. There are currently two kinds of plays that
// the players perform: tragedies and comedies. As well as providing a
// bill for the performance, the company gives is it's customers
// "volume credits", a loyalty mechanism they can use for discounts on
// future performances.

// The performers store data about their plays in a JSON file called 
// plays.json. They store data for their bills in a file called invoices.json
// The code that prints the bill is a function called statement.

// Running this code on the test data files results in the following output:
// Hamlet: USD 650.00 (55 seats)
// As You Like It: USD 580.00 (35 seats)
// Othello: USD 500.00 (40 seats)
// Amount owed is USD 1730.00
// You earned 47 credits

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
