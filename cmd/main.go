package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ropfoo/clobbopus/internal/config"
	"github.com/ropfoo/clobbopus/internal/download"
	"github.com/ropfoo/clobbopus/internal/helper"
)

func main() {
	fmt.Println("hello  🐙")
	config := config.Read()
	// create initial main pages folder
	os.Mkdir("../pages", os.ModePerm)

	for domain := range config.Domains {
		wg := sync.WaitGroup{}
		for _, query := range config.Domains[domain].Queries {
			url := config.Domains[domain].Url + "?" + query
			formattedQuery := strings.ReplaceAll(query, "/", "-")
			urlParts, pageRange := helper.ExtractPageRange(url)
			for page := pageRange.StartPage; page <= pageRange.EndPage; page++ {
				wg.Add(1)
				go download.DownloadPage(
					download.Page{
						Url:      urlParts[0] + fmt.Sprint(page) + urlParts[1],
						Filename: formattedQuery + fmt.Sprint(page),
						Dir:      "../pages/" + domain + "/",
					},
					&wg,
				)
			}
		}
		wg.Wait()
	}
}
