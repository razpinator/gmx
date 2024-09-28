package logic

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ConvertToKebabCase converts a string to kebab-case
func ConvertToKebabCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return strcase.ToKebab(str), nil
}

// ConvertToCamelCase converts a string to CamelCase
func ConvertToCamelCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return strcase.ToCamel(str), nil
}

// ConvertToSnakeCase converts a string to snake_case
func ConvertToSnakeCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return strcase.ToSnake(str), nil
}

// ConvertToPascalCase converts a string to PascalCase
func ConvertToPascaleCase(input interface{}, params ...interface{}) (interface{}, error) {
	str, ok := input.(string)
	if !ok {
		return nil, fmt.Errorf("expected string input")
	}
	return toPascalCase(str), nil
}

// toPascalCase converts a string to PascalCase
func toPascalCase(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", " ")
	parts := strings.Fields(s)
	caser := cases.Title(language.English)
	for i := range parts {
		parts[i] = caser.String(parts[i])
	}
	return strings.Join(parts, " ")
}
