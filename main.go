package main

import (
				"github.com/gofiber/fiber"
				"go-rest/film"	
				"go-rest/database"				
				"github.com/jinzhu/gorm"
				"os"
				"github.com/joho/godotenv"
				"fmt"
				_ "github.com/jinzhu/gorm/dialects/postgres"
				


)

var db *gorm.DB //database

func initDatabase(){
		var err error 
		e := godotenv.Load() //Load .env file
		if e != nil {
			fmt.Print(e)
		}


		dbURL := os.Getenv("db_url")
		fmt.Println(dbURL)

		database.DBConn, err =  gorm.Open("postgres", dbURL)
		if err != nil {
			fmt.Print(err)
			panic("failed to conenct ")
		}

		fmt.Println("Connection opened database")

		// database.DBConn.AutoMigrate(&film.Film{})/a

		fmt.Println("database migrated")

		// db = conn
		// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

func helloWorld( c *fiber.Ctx){
	c.Send("Hello world ")
}

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/film", film.GetFilms)
	app.Get("/api/v1/film/:id", film.GetFilm)
	app.Post("/api/v1/film", film.NewFilm)
	app.Delete("api/v1/film/:id", film.DelFilm)

	app.Get("/", helloWorld)

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)

	app.Listen("127.0.0.1:8080")

	defer database.DBConn.Close()

	

}
