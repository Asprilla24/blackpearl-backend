package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func anagramsGroup(arr []string) (result [][]string) {
	groups := make(map[string][]string)
	for _, val := range arr {
		sortedString := sortString(val)
		if v, ok := groups[sortedString]; ok {
			v = append(v, val)
			groups[sortedString] = v
		} else {
			groups[sortedString] = []string{val}
		}
	}

	for _, val := range groups {
		result = append(result, val)
	}

	return
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)

	return strings.Join(s, "")
}

func prettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

func main() {
	input := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	prettyPrint(anagramsGroup(input))
}
