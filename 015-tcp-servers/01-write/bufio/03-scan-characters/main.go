//scan a string type variable for tokens line and split by chars
//tokens by default are lines
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	s := "I felt so good like anything was possible\n I hit cruise control and rubbed my eyes\n"

	scanner := bufio.NewScanner(strings.NewReader(s))

	//runes = characters
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		// get the next token of text
		fmt.Printf("%s\n", scanner.Text())
	}

}
