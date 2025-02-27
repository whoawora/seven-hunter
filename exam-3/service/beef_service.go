package service

import (
	"io"
	"net/http"
	"strings"
)

func FetchAndCountBeef() (map[string]int, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	meatTypes := map[string]struct{}{
		"t-bone":   {},
		"fatback":  {},
		"pastrami": {},
		"pork":     {},
		"meatloaf": {},
		"jowl":     {},
		"enim":     {},
		"bresaola": {},
	}

	text := strings.ToLower(string(body))
	words := strings.FieldsFunc(text, func(r rune) bool {
		return r == ' ' || r == ',' || r == '.' || r == '\n' || r == '\t'
	})

	counts := make(map[string]int)
	for _, word := range words {
		word = strings.TrimSpace(word)
		if _, exists := meatTypes[word]; exists {
			counts[word]++
		}
	}
	for meat := range meatTypes {
		if _, exists := counts[meat]; !exists {
			counts[meat] = 0
		}
	}

	return counts, nil
}
