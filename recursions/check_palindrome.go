package main

import (
	"regexp"
	"strings"
)

var nonAlphaNumRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func isPalindrome(s string) bool {
	cleaned := nonAlphaNumRegex.ReplaceAllString(s, "")
	lowerStr := strings.ToLower(cleaned)
	return checkPalindrom(lowerStr, 0, len(lowerStr)-1)
}

func checkPalindrom(s string, l int, r int) bool {
	if l >= r {
		return true
	}

	if s[l] != s[r] {
		return false
	}

	return checkPalindrom(s, l+1, r-1)
}
