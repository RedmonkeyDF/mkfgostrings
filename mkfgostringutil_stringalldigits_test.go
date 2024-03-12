package mfstringutil

import (
	"strings"
	"testing"
)

func TestStringAllDigits(t *testing.T) {

	tst := make(map[string]bool)

	tst["1234567890"] = true
	tst["12.23"] = true
	tst["324534.44.44"] = false
	tst[""] = false
	tst["238 2387"] = false
	tst["2346236."] = true

	for inval, expected := range tst {

		got := StringAllDigits(inval, strings.Count(inval, ".") > 0)

		if got != expected {

			t.Fatalf("%s failed.  Input \"%s\".  Expected \"%t\".", t.Name(), inval, expected)
		}
	}
}

func TestStringAllZeros(t *testing.T) {

	tst := make(map[string]bool)

	tst["1234567890"] = false
	tst["0"] = true
	tst["."] = true
	tst[""] = false
	tst[".0"] = true
	tst["0.00"] = true
	tst["12."] = false
	tst["00000.0"] = true
	tst["0."] = true

	for inval, expected := range tst {

		got := StringAllZeros(inval, strings.Count(inval, ".") > 0)

		if got != expected {

			t.Fatalf("%s failed.  Input \"%s\".  Expected \"%t\".", t.Name(), inval, expected)
		}
	}

}