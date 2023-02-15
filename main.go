package main // editing /auto correction tool

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func stringremove(slice []string, s int) []string { // Removes a string from a slice of string
	return append(slice[:s], slice[s+1:]...)
}

func runeremove(slice []rune, s int) []rune { // Removes a string from a slice of string
	return append(slice[:s], slice[s+1:]...)
}

func main() { // open and scan
	var strarr []string
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)
	for fileScanner.Scan() {
		strarr = append(strarr, fileScanner.Text())
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	} // calling in functions
	for i, str := range strarr {
		// fmt.Println(i, len(str), str)
		if str == "(hex)" {
			strarr[i-1] = hexaNumberToInteger(strarr[i-1])
			strarr = remove(strarr, i)
			i--
		}

		if str == "(bin)" {
			strarr[i-1] = bintodecimal(strarr[i-1])
			strarr = remove(strarr, i)
			i--
		}

		if str == "(up)" {
			strarr[i-1] = toUpper(strarr[i-1])
			strarr = remove(strarr, i)
			i--
		}

		if str == "(low)" {
			strarr[i-1] = tolower(strarr[i-1])
			strarr = remove(strarr, i)
			i--
		}

		if str == "(cap)" {
			strarr[i-1] = caps(strarr[i-1])
			strarr = remove(strarr, i)
			i--
		}

		// (remove hex/bin etc)
		if str == "(cap," {
			x := trimatoi(strarr[i+1])
			for u := 1; u <= x; u++ {
				strarr[i-u] = caps(strarr[i-u])
			}
			strarr = remove(strarr, i)
			i--
			strarr = remove(strarr, i+1)
			i--
		} // (remove up+number),

		if str == "(up," {
			x := trimatoi(strarr[i+1])
			for u := 1; u <= x; u++ {
				strarr[i-u] = toUpper(strarr[i-u])
			}
			strarr = remove(strarr, i)
			i--
			strarr = remove(strarr, i+1)
			i--
		}

		if str == "(low," { // remove low
			x := trimatoi(strarr[i+1])
			for u := 1; u <= x; u++ {
				strarr[i-u] = tolower(strarr[i-u])
			}
			strarr = remove(strarr, i)
			i--
			strarr = remove(strarr, i+1)
			i--
		}
	} // output amends to result file
	strarr = punct(strarr)
	strarr = speech(strarr)
	strarr = atoan(strarr)
	resultString := strings.Join(strarr, " ")
	file, err = os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(resultString)
	file.Sync()
}

// trim and atoi function
func trimatoi(s string) int {
	srune := []rune(s)
	n := 0
	for _, rune := range srune {
		if rune >= '0' && rune <= '9' {
			y := int(rune - '0')
			n = n*10 + y
		}
	}
	return n
}

func getAlpha(srune []rune) string { // good luck code
	str := ""
	for _, rune := range srune {
		if rune >= 'a' || rune <= 'z' || rune >= 'A' || rune <= 'Z' {
			str += string(rune)
		}
	}
	return str
}

func hexaNumberToInteger(hexaString string) string {
	i, _ := strconv.ParseInt(hexaString, 16, 64)
	return fmt.Sprint(i)
}

//(bin should replace the word before with the decimal version of the word (in this case the word will always be a binary number).

func bintodecimal(binToString string) string {
	i, _ := strconv.ParseInt(binToString, 2, 64)
	return fmt.Sprint(i)
}

// Every instance of (up) converts the word placed before in the Uppercase version of it.

func toUpper(str string) string {
	return strings.ToUpper(str)
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

// Every instance of (low) converts the word placed before in the Lowercase version of it.

func tolower(str string) string {
	return strings.ToLower(str)
}

// Every instance of multiples on punctuations like !!/... -
func punct(strarr []string) []string {
	for i := 0; i < len(strarr); i++ {
		count := 0
		found := true
		for u, strRune := range strarr[i] {
			for x, multiRune := range strarr[i] {
				if (multiRune == '.' || multiRune == ',' || multiRune == '!' || multiRune == '?' || multiRune == ':' || multiRune == ';') && x == 0 {
					found = false
				}
				if (multiRune == '.' || multiRune == ',' || multiRune == '!' || multiRune == '?' || multiRune == ':' || multiRune == ';') && !found {
					count++
				}
				if (multiRune == '.' || multiRune == ',' || multiRune == '!' || multiRune == '?' || multiRune == ':' || multiRune == ';') && count == len(strarr[i])-1 {
					strarr[i-1] += strarr[i]
					strarr = remove(strarr, i)
					i--
					found = true
				}
			} // every instance of singles on punctuations - '.', ',', '!', '?', ':' and ';' should be close to the previous word and with space apart from the next one.
			if (strRune == '.' || strRune == ',' || strRune == '!' || strRune == '?' || strRune == ':' || strRune == ';') && u == 0 && !found {
				if len(strarr[i]) == 1 && u == 0 {
					strarr[i-1] += string(strRune)
					strarr = remove(strarr, i)
					i--
				} else if len(strarr[i]) > 1 && u == 0 && !found {
					strarr[i-1] += string(strRune)
					srune := []rune(strarr[i])
					srune = runeremove(srune, u)
					strarr[i] = string(srune)
				}
			}
		}
	}
	return strarr
}

// The punctuation mark ''' should not have spaces if there are letters in both sides of it. Otherwise, the mark should be placed to the right of the next word and the following ''' mark should be placed to its left.

func speech(strarr []string) []string {
	quoteFound := false
	for i := 0; i < len(strarr); i++ {
		if strarr[i] == "'" {
			if !quoteFound {
				strarr[i+1] = "'" + strarr[i+1]
				strarr = remove(strarr, i)
				i--
				quoteFound = true
			} else {
				strarr[i-1] += "'"
				strarr = remove(strarr, i)
				i--
			}
		}
	}
	return strarr
}

// removing instance Every instance of (cap) converts the letter placed before into a capital.
func caps(s string) string {
	s = tolower(s)
	srune := []rune(s)
	srune[0] = srune[0] - 32
	return string(srune)
}

// replacing every instance of "a" turned into "an" if the next words begins with a vowel or an 'h'

func atoan(strarr []string) []string {
	for i, srune := range strarr {
		for _, rune := range srune {
			if (rune == 'a' || rune == 'A') && len(srune) == 1 {
				if strarr[i+1][0] == 'a' || strarr[i+1][0] == 'e' ||
					strarr[i+1][0] == 'i' ||
					strarr[i+1][0] == 'o' ||
					strarr[i+1][0] == 'u' ||
					strarr[i+1][0] == 'h' {
					strarr[i] += "n"
				}
			}
		}
	}
	return strarr
}
