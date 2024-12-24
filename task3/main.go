package main

import (
	"fmt"
	"strconv"
	"strings"
)

func UnboxingFunction(s string) string {
	if len(s) == 0 || strings.Contains("0123456789", string(s[0])) {
		return ""
	}

	sb := strings.Builder{}
	for i, c := range s {
		if digit, err := strconv.Atoi(string(c)); err == nil {
			for j := 1; j < digit; j++ {
				sb.WriteString(string(s[i-1]))
			}
		} else {
			sb.WriteString(string(c))
		}
	}

	return sb.String()
}

func main() {
	fmt.Println(UnboxingFunction("a4bc2d5e"))
}
