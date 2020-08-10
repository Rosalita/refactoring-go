// The first step in factoring is always the same, ensure a solid set of tests
// exist for that section of code. These tests must be self-checking.

package main

import (
	"errors"
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	flag.Parse()
	// Test setup
	playFor = mockPlayFor

	// Run tests
	exitCode := m.Run()

	// Test teardown
	os.Exit(exitCode)
}

func mockPlayFor(perf Performance) Play {
	switch perf.PlayID {
	case "shrew":
		return Play{Name: "The Taming of the Shrew", Type: "comedy"}
	case "a-and-c":
		return Play{Name: "Antony and Cleopatra", Type: "tragedy"}
	case "foo":
		return Play{Name: "Foo Bar", Type: "baz"}
	}

	return Play{}
}

func TestStatement(t *testing.T) {

	comedySmallInvoice := Invoice{
		Customer:     "Rosie",
		Performances: []Performance{{PlayID: "shrew", Audience: 20}},
	}

	comedyLargeInvoice := Invoice{
		Customer:     "Rosie",
		Performances: []Performance{{PlayID: "shrew", Audience: 200}},
	}

	tragedySmallInvoice := Invoice{
		Customer:     "Rosie",
		Performances: []Performance{{PlayID: "a-and-c", Audience: 20}},
	}

	tragedyLargeInvoice := Invoice{
		Customer:     "Rosie",
		Performances: []Performance{{PlayID: "a-and-c", Audience: 200}},
	}

	unknownInvoice := Invoice{
		Customer:     "Rosie",
		Performances: []Performance{{PlayID: "foo", Audience: 200}},
	}

	tests := []struct {
		invoice        Invoice
		expectedResult string
		err            error
	}{
		{comedySmallInvoice,
			"Statement for Rosie \n" +
				"The Taming of the Shrew: USD 360.00 (20 seats) \n" +
				"Amount owed is USD 360.00\nYou earned 4 credits\n",
			nil},
		{comedyLargeInvoice,
			"Statement for Rosie \n" +
				"The Taming of the Shrew: USD 1900.00 (200 seats) \n" +
				"Amount owed is USD 1900.00\nYou earned 210 credits\n",
			nil},
		{tragedySmallInvoice,
			"Statement for Rosie \n" +
				"Antony and Cleopatra: USD 400.00 (20 seats) \n" +
				"Amount owed is USD 400.00\nYou earned 0 credits\n",
			nil},
		{tragedyLargeInvoice,
			"Statement for Rosie \n" +
				"Antony and Cleopatra: USD 2100.00 (200 seats) \n" +
				"Amount owed is USD 2100.00\nYou earned 170 credits\n",
			nil},
		{unknownInvoice, "", errors.New("error: unknown performance type baz")},
	}
	for _, test := range tests {
		result, err := statement(test.invoice)
		assert.Equal(t, test.expectedResult, result)
		assert.Equal(t, test.err, err)
	}
}
