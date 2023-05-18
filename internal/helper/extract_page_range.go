package helper

import (
	"strconv"
	"strings"
)

type PageRange struct {
	StartPage int
	EndPage   int
}

// Extracts page count from url
//
// ~1-5~ -> {StartPage: 1, EndPage: 5}
func ExtractPageRange(url string) ([2]string, PageRange) {
	var pageRange PageRange
	var urlParts [2]string = [2]string{url, ""}
	if strings.Contains(url, "~") {
		splittedUrl := strings.Split(url, "~")
		urlParts = [2]string{splittedUrl[0], splittedUrl[2]}
		pageRangePart := splittedUrl[1]
		// get start and end page
		pages := strings.Split(pageRangePart, "-")
		startPage, error := strconv.Atoi(pages[0])
		if error != nil {
			startPage = 1
		}
		endPage, error := strconv.Atoi(pages[1])
		if error != nil {
			endPage = 1
		}
		pageRange = PageRange{
			StartPage: startPage,
			EndPage:   endPage,
		}
	}
	return urlParts, pageRange
}
