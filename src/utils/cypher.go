package utils

import (
	"fmt"
	"strings"
)

func ListToCypher(list []string) string {
	var quoted []string
	for _, value := range list {
		quoted = append(quoted, fmt.Sprintf(`"%s"`, value))
	}

	return strings.Join(quoted, ", ")
}
