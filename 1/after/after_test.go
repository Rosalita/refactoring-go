// The first step in factoring is always the same, ensure a solid set of tests
// exist for that section of code. These tests must be self-checking.

package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatement(t *testing.T) {

	comedyPlays := map[string]Play{
		"shrew": {Name: "The Taming of the Shrew", Type: "comedy"},
	}

	tragedyPlays := map[string]Play{
		"a-and-c": {Name: "Antony and Cleopatra", Type: "tragedy"},
	}

	unknownPlays := map[string]Play{
		"foo": {Name: "Foo Bar", Type: "baz"},
	}

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
		plays          map[string]Play
		expectedResult string
		err            error
	}{
		{comedySmallInvoice, comedyPlays,
			"Statement for Rosie \n" +
				"The Taming of the Shrew: USD 360.00 (20 seats) \n" +
				"Amount owed is USD 360.00\nYou earned 4 credits\n",
			nil},
		{comedyLargeInvoice, comedyPlays,
			"Statement for Rosie \n" +
				"The Taming of the Shrew: USD 1900.00 (200 seats) \n" +
				"Amount owed is USD 1900.00\nYou earned 210 credits\n",
			nil},
		{tragedySmallInvoice, tragedyPlays,
			"Statement for Rosie \n" +
				"Antony and Cleopatra: USD 400.00 (20 seats) \n" +
				"Amount owed is USD 400.00\nYou earned 0 credits\n",
			nil},
		{tragedyLargeInvoice, tragedyPlays,
			"Statement for Rosie \n" +
				"Antony and Cleopatra: USD 2100.00 (200 seats) \n" +
				"Amount owed is USD 2100.00\nYou earned 170 credits\n",
			nil},
		{unknownInvoice, unknownPlays, "", errors.New("error: unknown performance type baz")},
	}
	for _, test := range tests {
		result, err := statement(test.invoice, test.plays)
		assert.Equal(t, test.expectedResult, result)
		assert.Equal(t, test.err, err)
	}
}
