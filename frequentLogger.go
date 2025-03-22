package main

import "strings"

func buildEventsString(events []string) string {
	var builder strings.Builder
	builder.WriteString("frequent events:\n")

	for _, event := range events {
		builder.WriteString(event)
		builder.WriteRune('\n')
	}
	return builder.String()
}
