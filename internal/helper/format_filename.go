package helper

import (
	"strings"
)

func ConvertToFilename(input string) string {
	var filename string
	filename = strings.SplitN(input, "/", 4)[3]
	// replace slash
	filename = strings.ReplaceAll(filename, "/", "~")
	// replace query start
	filename = strings.ReplaceAll(filename, "?", "QUERY")
	// replace AND
	filename = strings.ReplaceAll(filename, "&", "AND")
	return filename
}

func RevertFilename(input string) string {
	var filename string
	// replace slash
	filename = strings.ReplaceAll(input, "~", "/")
	// replace query start
	filename = strings.ReplaceAll(filename, "QUERY", "?")
	// replace AND
	filename = strings.ReplaceAll(filename, "AND", "&")
	filename = strings.ReplaceAll(filename, ".html", "")
	return filename
}
