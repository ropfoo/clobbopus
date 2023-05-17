package download

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

func DownloadPage(url string, filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)

	// request page
	response, error := http.Get(url)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer response.Body.Close()

	os.Mkdir("../test", os.ModePerm)

	f := fmt.Sprintf("../test/%q.html", filename)
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
