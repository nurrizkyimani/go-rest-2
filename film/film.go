package film

import (
	"github.com/gofiber/fiber"
)

func GetFilms(c *fiber.Ctx) {
	c.Send("Send Films this is awesome")
}

func GetFilm(c *fiber.Ctx) {
	c.Send(" Get Film ")
}

func NewFilm(c *fiber.Ctx) {
	c.Send("Post new Film")
}

func UpdateFilm(c *fiber.Ctx) {
	c.Send("Send Films")
}

func DelFilm(c *fiber.Ctx) {
	c.Send("Send Films")
}
