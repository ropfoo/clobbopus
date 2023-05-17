package main

import (
	"fmt"
	"sync"

	"github.com/ropfoo/clobbopus/internal/config"
	"github.com/ropfoo/clobbopus/internal/download"
)

func main() {
	fmt.Println("hello  🐙")

	config := config.Read()

	for domain := range config.Domains {
		wg := sync.WaitGroup{}
		for index, url := range config.Domains[domain].Urls {
			wg.Add(1)
			go download.DownloadPage(url, domain+fmt.Sprint(index), &wg)
		}
		wg.Wait()
	}

}
