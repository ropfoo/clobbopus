package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

type Page struct {
	Url      string
	Filename string
	Dir      string
}

func DownloadPage(page Page, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)

	// request page
	response, error := http.Get(page.Url)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer response.Body.Close()

	os.Mkdir(page.Dir, os.ModePerm)
	f := page.Dir + page.Filename + ".html"
	// create file
	file, error := os.Create(f)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer file.Close()

	// copy html into file
	_, error = io.Copy(file, response.Body)
	if error != nil {
		fmt.Println(error)
		return
	}
}
