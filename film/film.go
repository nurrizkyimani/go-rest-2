package film

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/satori/go.uuid"
	"go-rest/database"
)

type Film struct {
	gorm.Model 
	ID   uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title  string `json:"title"`
	Director string `json:"director"` 
	createdAt time.Time 

}

func GetFilms(c *fiber.Ctx) {
	c.Send("Send Films this is awesome")
	db := database.DBConn
	var films []Film
	db.Find(&films)
	c.JSON(films)
}

func GetFilm(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var films []Film
	db.Find(&films, id)
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
