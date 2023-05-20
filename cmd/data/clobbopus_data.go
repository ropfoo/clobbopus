package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/ropfoo/clobbopus/internal/config"
	"github.com/ropfoo/clobbopus/internal/download"
	"github.com/ropfoo/clobbopus/internal/helper"
)

func main() {
	fmt.Println("🐙 ~~~downloading files~~~")
	config := config.Read()
	// create initial main pages folder
	os.Mkdir(config.Dist, os.ModePerm)

	for domain := range config.Domains {
		wg := sync.WaitGroup{}
		for _, params := range config.Domains[domain].Params {
			var domainUrl = config.Domains[domain].Url
			paramsParts, pageRange := helper.ExtractPageRangeFromParams(params)
			for page := pageRange.StartPage; page <= pageRange.EndPage; page++ {
				wg.Add(1)
				// create url to request page from
				url := domainUrl + paramsParts[0] + fmt.Sprint(page) + paramsParts[1]
				fmt.Println("URL: ", url)
				filename := helper.ConvertToFilename(url)
				go download.DownloadPage(
					download.Page{
						Url:      url,
						Filename: filename,
						Dir:      config.Dist + "/" + domain + "/",
					},
					&wg,
				)
			}
		}
		wg.Wait()
	}
}

// http://localhost:3000/musicstuff_/page=5?instrument=bass
