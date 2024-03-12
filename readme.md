# MKFGOSTRINGS

Simple string utilities.  I wrote these as I was tired of typing the same code into my projects just to check 
if a string consists of all digits.

##### v1.0.0

`func StringAllDigits(instr string, acceptperiod bool) bool`

Determines whether a string is made up of all digits.  Optionally specifying if it accepts a period character.

`func StringAllZeros(instr string, acceptperiod bool) bool`

Determines whether a string is made up of all zeros.  Optionally specifying if it accepts a period character.

`func FormatDuration(dtn time.Duration) string`

Formats a duration with a slightly more pleasing (to me) format.