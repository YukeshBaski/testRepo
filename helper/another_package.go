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
