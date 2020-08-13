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
	PlayID        string `json:"playID"`
	Audience      int    `json:"audience"`
	Play          Play  
	Amount        int
	VolumeCredits int
}

type Play struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type StatementData struct {
	Customer           string
	Performances       []Performance
	TotalAmount        float64
	TotalVolumeCredits int
}

var playFor func(Performance) Play

func main() {

	// Global variable used for test mock injection
	playFor = playForFunc

	invoiceFile, err := ioutil.ReadFile("invoices.json")
	if err != nil {
		fmt.Println(err)
	}

	var invoice Invoice
	if err := json.Unmarshal(invoiceFile, &invoice); err != nil {
		fmt.Println(err)
	}

	result, err := statement(invoice)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func renderPlainText(data StatementData) (string, error) {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("Statement for %s \n", data.Customer))

	for _, perf := range data.Performances {
		result.WriteString(fmt.Sprintf("%s: %s (%d seats) \n", perf.Play.Name, usd(float64(perf.Amount)), perf.Audience))
	}
	result.WriteString(fmt.Sprintf("Amount owed is %s\n", usd(data.TotalAmount)))
	result.WriteString(fmt.Sprintf("You earned %d credits\n", data.TotalVolumeCredits))
	return result.String(), nil
}

func statement(invoice Invoice) (string, error) {
	var enrichedPerformances []Performance
	for _, perf := range invoice.Performances {
		enrichedPerf, err := enrichPerformance(perf)
		if err != nil {
			return "", err
		}
		enrichedPerformances = append(enrichedPerformances, *enrichedPerf)
	}

	statementData := StatementData{
		Customer:           invoice.Customer,
		Performances:       enrichedPerformances,
		TotalAmount:        totalAmount(enrichedPerformances),
		TotalVolumeCredits: totalVolumeCredits(enrichedPerformances),
	}
	return renderPlainText(statementData)
}

func enrichPerformance(perf Performance) (*Performance, error) {
	perf.Play = playFor(perf)
	amount, err := amountFor(perf)
	if err != nil {
		return nil, err
	}
	volumeCredits := volumeCreditsFor(perf)

	enrichedPerf := Performance{
		PlayID:        perf.PlayID,
		Audience:      perf.Audience,
		Play:          perf.Play,
		Amount:        amount,
		VolumeCredits: volumeCredits,
	}

	return &enrichedPerf, nil
}

func totalAmount(performances []Performance) float64 {
	var result float64
	for _, perf := range performances {
		amount, err := amountFor(perf)
		if err != nil {
			fmt.Println(err)
		}
		result += float64(amount)
	}
	return result
}

func totalVolumeCredits(performances []Performance) int {
	var result int
	for _, perf := range performances {
		result += perf.VolumeCredits
	}
	return result
}

func usd(amount float64) string {
	return fmt.Sprintf("%+v", currency.USD.Amount(amount/100))
}

func volumeCreditsFor(perf Performance) int {
	var result int
	result += int(math.Max(float64(perf.Audience)-30, 0))
	if "comedy" == playFor(perf).Type {
		result += int(math.Floor(float64(perf.Audience) / 5))
	}
	return result
}

func playForFunc(perf Performance) Play {
	playsFile, err := ioutil.ReadFile("plays.json")
	if err != nil {
		fmt.Println(err)
	}

	var plays map[string]Play
	if err := json.Unmarshal(playsFile, &plays); err != nil {
		fmt.Println(err)
	}
	return plays[perf.PlayID]
}

func amountFor(perf Performance) (int, error) {
	var result int

	switch perf.Play.Type {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += 1000 * (perf.Audience - 30)
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += 10000 + 500*(perf.Audience-20)
		}
		result += 300 * perf.Audience

	default:
		return result, fmt.Errorf("error: unknown performance type %s", perf.Play.Type)
	}
	return result, nil
}
