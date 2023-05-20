package helper

import (
	"strconv"
	"strings"
)

type PageRange struct {
	StartPage int
	EndPage   int
}

const PAGE_RANGE_SPLITTER string = "~"

// Extracts page count from url
//
// ~1-5~ -> {StartPage: 1, EndPage: 5}
func ExtractPageRangeFromParams(params string) ([2]string, PageRange) {
	var pageRange PageRange
	var paramsParts [2]string = [2]string{params, ""}
	if strings.Contains(params, PAGE_RANGE_SPLITTER) {
		// split params into 3 parts (start of qparamsuery, page range, end of params)
		splittedParams := strings.Split(params, PAGE_RANGE_SPLITTER)

		// params without page range part
		paramsParts = [2]string{splittedParams[0], splittedParams[2]}

		pageRangePart := splittedParams[1]

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
	return paramsParts, pageRange
}
