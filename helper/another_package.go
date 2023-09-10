package helper

import (
	"strings"
)

func HelperFunc() string {
	return "Im a helper"
}

// from login branch
func IsPalindrome(str string) bool {
	sb := strings.Builder{}

	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteString(string(str[i]))
	}
	if sb.String() != str {
		return false
	}
	return true
}

// from new branch optimized solution

func IsPalyndromeFromNewBranch(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
