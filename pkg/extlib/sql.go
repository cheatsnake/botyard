package extlib

import "strings"

func SQLQueryPlaceholders(length int) string {
	if length < 1 {
		return ""
	}

	placeholders := make([]string, length)
	for i := range placeholders {
		placeholders[i] = "?"
	}

	return strings.Join(placeholders, ",")
}
