package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	//
	htmlFolders, err := os.ReadDir("../../pages")
	if err != nil {
		panic("oh no")
	}
	//
	for _, htmlFolder := range htmlFolders {
		if htmlFolder.IsDir() {
			domain := htmlFolder.Name()
			htmlFiles, err := os.ReadDir("../../pages/" + domain)
			if err != nil {
				panic("oh no")
			}
			//
			for _, htmlFile := range htmlFiles {
				name := htmlFile.Name()
				var query string
				// route := helper.RevertFilename(name)
				query = strings.ReplaceAll(name, ".html", "")
				query = strings.ReplaceAll(query, "~", "/")
				query = strings.ReplaceAll(query, "AND", "&")
				// query = strings.ReplaceAll(query, "QUERY", "?")
				route := "/" + domain + "/" + query
				fmt.Println("route:", route)
				app.Get(route, func(c *fiber.Ctx) error {
					return c.SendFile("../../pages/" + domain + "/" + name)
				})
			}
		}
	}
	log.Fatal(app.Listen(":3000"))
}
