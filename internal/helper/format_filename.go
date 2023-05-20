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

func ConvertUrlToFilename(url string, fileType string) string {
	var filename string
	// replace slash
	filename = strings.ReplaceAll(url, "/", "~")
	// replace query start
	filename = strings.ReplaceAll(filename, "?", "QUERY")
	// replace AND
	filename = strings.ReplaceAll(filename, "&", "AND")
	// add file type
	filename = filename + fileType
	return filename
}
