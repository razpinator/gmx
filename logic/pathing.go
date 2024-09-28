package logic

import (
	"fmt"
	"path/filepath"
)

// JoinPath joins strings into a path using the OS separator
func JoinPath(input interface{}, params ...interface{}) (interface{}, error) {
	parts, ok := input.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected input to be a list of strings")
	}

	strParts := make([]string, len(parts))
	for i, part := range parts {
		strPart, ok := part.(string)
		if !ok {
			return nil, fmt.Errorf("expected all elements to be strings")
		}
		strParts[i] = strPart
	}

	return filepath.Join(strParts...), nil
}
