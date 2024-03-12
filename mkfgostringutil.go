package mfstringutil

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func StringAllDigits(instr string, acceptperiod bool) bool {

	if utf8.RuneCountInString(instr) == 0 {

		return false
	}

	isNotDigit := func(c rune) bool {

		isnotadigit := c < '0' || c > '9'

		if acceptperiod {

			if isnotadigit && (c == '.') {

				isnotadigit = false
			}
		}

		return isnotadigit
	}

	alldigits := strings.IndexFunc(instr, isNotDigit) == -1

	if acceptperiod && alldigits {

		alldigits = strings.Count(instr, ".") < 2
	}

	return alldigits
}

func StringAllZeros(instr string, acceptperiod bool) bool {

	if utf8.RuneCountInString(instr) == 0 {

		return false
	}

	isNotZero := func(c rune) bool {

		isnotazero := c != '0'

		if acceptperiod {

			if isnotazero && (c == '.') {

				isnotazero = false
			}
		}

		return isnotazero
	}

	allzeros := strings.IndexFunc(instr, isNotZero) == -1

	if acceptperiod && allzeros {

		allzeros = strings.Count(instr, ".") < 2
	}

	return allzeros
}

func FormatDuration(dtn time.Duration) string {

	outstr := ""
	addout := func(num, pad int, numtag string) {

		if utf8.RuneCountInString(outstr) > 0 && num > 0 {

			outstr += " "
		}

		if num > 0 {

			outstr += fmt.Sprintf("%."+strconv.Itoa(pad)+"d %s", num, numtag)
		}
	}

	d := dtn.Round(time.Millisecond)
	days := d / (time.Hour * 24)
	d -= days * (time.Hour * 24)
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second
	d -= seconds * time.Second
	mseconds := d / time.Millisecond

	addout(int(days), 0, "day")
	if int(days) > 1 {
		outstr += "s"
	}
	addout(int(hours), 0, "hr")
	addout(int(minutes), 2, "min")
	addout(int(seconds), 2, "sec")
	addout(int(mseconds), 3, "ms")

	return outstr
}
