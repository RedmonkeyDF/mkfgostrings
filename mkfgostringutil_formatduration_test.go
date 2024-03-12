package mfstringutil

import (
	"fmt"
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {

	type TestObj struct {

		duration time.Duration
		expected string
	}

	tm := time.Date(2000, 0, 0, 0, 0, 0, 0, time.UTC)

	testcases := []TestObj{

		{tm.Sub(time.Date(2000, 0, 0, -1, -1, -1, int(-12*time.Millisecond), time.UTC)), "1 hr 01 min 01 sec 012 ms"},
		{tm.Sub(time.Date(2000, 0, 0, -1, 0, 0, 0, time.UTC)), "1 hr"},
		{tm.Sub(time.Date(2000, 0, 0, 0, -23, -5, 0, time.UTC)), "23 min 05 sec"},
		{tm.Sub(time.Date(2000, 0, -12, 0, -23, 0, int(-234*time.Millisecond), time.UTC)), "12 days 23 min 234 ms"},
		{tm.Sub(time.Date(2000, 0, -1, 0, -23, 0, int(-234*time.Millisecond), time.UTC)), "1 day 23 min 234 ms"},
	}

	for idx, tcase := range testcases {

		t.Run(fmt.Sprintf("testcase:%.2d", idx+1), func(t *testing.T) {

			got := FormatDuration(tcase.duration)

			if got != tcase.expected {

				t.Fatalf("Expected: \"%s\".  Got: \"%s\".", tcase.expected, got)
			}
		})
	}


}
