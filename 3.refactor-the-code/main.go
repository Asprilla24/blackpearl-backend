package main

import (
	"fmt"
	"strings"
)

func findFirstStringInBracket(str string) string {
	firstBracketIndex := strings.Index(str, "(")
	if firstBracketIndex >= 0 {
		closingBracketIndex := strings.Index(str, ")")
		if closingBracketIndex >= 0 {
			return str[firstBracketIndex+1 : closingBracketIndex]
		}
	}

	return ""
}

func main() {
	fmt.Println(findFirstStringInBracket("asdas(a)"))
}
