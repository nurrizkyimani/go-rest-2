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
	Rating int `json:"rating"`

}

func GetFilms(c *fiber.Ctx) {
	// c.Send("Send Films this is awesome")
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
	db := database.DBConn


	film := new(Film)
	if err := c.BodyParser(film); err != nil {
		c.Status(503).Send(err)
		return 
	}

	db.Create(&film)
	c.JSON(film)
}

func UpdateFilm(c *fiber.Ctx) {
	c.Send("Send Films")
}

func DelFilm(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var film Film
	db.First(&film, id)

	if film.Title == "" {
		c.Status(500).Send("No film on given id")
		return
	}

	db.Delete(&film)

	c.Send("Successfluu")


}
