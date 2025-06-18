package main

import "strings"

func join(events []string, separator string) string {
	var builder strings.Builder

	for _, event := range events {
		builder.WriteString(event)
		builder.WriteString(separator)
	}
	return builder.String()
}
