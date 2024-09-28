package logic

import (
	"fmt"
	"unicode"
)

// LowerFirst converts the first letter of a string to lower case
func LowerFirst(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	if str == "" {
		return str, nil
	}
	runes := []rune(str)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes), nil
}
