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
	os.Mkdir("../../pages", os.ModePerm)

	for domain := range config.Domains {
		wg := sync.WaitGroup{}
		for _, query := range config.Domains[domain].Queries {
			configUrl := config.Domains[domain].Url + "?" + query
			urlParts, pageRange := helper.ExtractPageRange(configUrl)
			for page := pageRange.StartPage; page <= pageRange.EndPage; page++ {
				wg.Add(1)
				url := urlParts[0] + fmt.Sprint(page) + urlParts[1]
				filename := helper.ConvertToFilename(url)
				go download.DownloadPage(
					download.Page{
						Url:      url,
						Filename: filename,
						Dir:      "../../pages/" + domain + "/",
					},
					&wg,
				)
			}
		}
		wg.Wait()
	}
}
