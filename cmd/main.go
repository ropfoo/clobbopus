package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/ropfoo/clobbopus/internal/config"
	"github.com/ropfoo/clobbopus/internal/download"
)

func main() {
	fmt.Println("hello  🐙")

	config := config.Read()

	// create initial main pages folder
	os.Mkdir("../pages", os.ModePerm)

	for domain := range config.Domains {
		wg := sync.WaitGroup{}
		for _, query := range config.Domains[domain].Queries {
			wg.Add(1)
			url := config.Domains[domain].Url + "?" + query
			formattedQuery := strings.ReplaceAll(query, "/", "-")
			go download.DownloadPage(
				download.Page{
					Url:      url,
					Filename: formattedQuery,
					Dir:      "../pages/" + domain + "/",
				},
				&wg,
			)
		}
		wg.Wait()
	}

}
