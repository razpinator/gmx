package logic

import (
	"fmt"

	"github.com/gertd/go-pluralize"
)

func Pluralize(input interface{}, params ...interface{}) (interface{}, error) {
	word, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	pluralizer := pluralize.NewClient()
	return pluralizer.Plural(word), nil
}
