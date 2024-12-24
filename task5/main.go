package main

import (
	"fmt"
	"sort"
	"strings"
)

func search(words *[]string) *map[string][]string {
	result := make(map[string][]string)
	toLower(*words)
	sort.Strings(*words)

	for _, w := range *words {
		isStored := false
		for k, v := range result {
			if isAnagram(k, w) {
				result[k] = append(v, w)
				isStored = true
				break
			}
		}
		if !isStored {
			result[w] = append(result[w], w)
		}
	}

	for k, v := range result {
		if len(v) <= 1 {
			delete(result, k)
		}
	}

	return &result
}

func isAnagram(first, second string) bool {
	return sortChars(first) == sortChars(second)
}

func sortChars(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

func toLower(data []string) {
	for i, x := range data {
		data[i] = strings.ToLower(x)
	}
}

func main() {
	words := []string{"Тяпка", "пятак", "лИсток", "столик", "пятка", "слиток"}
	m := search(&words)
	fmt.Println(m)
}
