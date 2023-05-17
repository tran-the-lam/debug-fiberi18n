package main

import (
	"test/fiberi18n"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/language"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()
	app.Use(
		fiberi18n.New(&fiberi18n.Config{
			RootPath:        "./localize",
			AcceptLanguages: []language.Tag{language.English},
			DefaultLanguage: language.English,
		}),
	)

	app.Get("/", func(c *fiber.Ctx) error {
		log.Print("HIHI")
		return c.SendString(fiberi18n.MustGetMessage("welcome"))
	})

	app.Listen("127.0.0.1:3000")
}
