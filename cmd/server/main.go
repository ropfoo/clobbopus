package main

import (
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ropfoo/clobbopus/internal/config"
	"github.com/ropfoo/clobbopus/internal/helper"
)

func main() {
	app := fiber.New()
	config := config.Read()
	//
	htmlFolders, err := os.ReadDir("../../pages")
	if err != nil {
		panic("oh no")
	}
	for _, htmlFolder := range htmlFolders {
		if htmlFolder.IsDir() {
			var domain = htmlFolder.Name()
			var route = config.Route(domain)
			var subRoute = config.Subroute(domain)

			app.Get(route+"/*", func(c *fiber.Ctx) error {
				var uri string = c.Context().URI().String()
				var uriHasSlashParam bool = strings.Contains(uri, "/?")
				var params string = c.Params("*")
				var queryString string = c.Context().QueryArgs().String()
				var url string = subRoute
				if params != "" {
					url = url + "/" + params
				}
				if queryString != "" {
					if uriHasSlashParam {
						url = url + "/" + "?" + queryString
					} else {
						url = url + "?" + queryString
					}
				}
				var filename string = helper.ConvertUrlToFilename(url, ".html")
				return c.SendFile("../../pages/" + domain + "/" + filename)
			})
		}
	}
	log.Fatal(app.Listen(":3000"))
}
