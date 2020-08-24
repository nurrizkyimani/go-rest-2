package main

import ("github.com/gofiber/fiber"
				"go-rest/film"	
				// "go-rest/database"				
				"github.com/jinzhu/gorm"
				"os"
				"github.com/joho/godotenv"
				"fmt"


)

var db *gorm.DB //database

func initDatabase(){
		e := godotenv.Load() //Load .env file
		if e != nil {
			fmt.Print(e)
		}

		username := os.Getenv("db_user")
		password := os.Getenv("db_pass")
		dbName := os.Getenv("db_name")
		dbHost := os.Getenv("db_host")
		dbPort := os.Getenv("db_port")


		dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password) //Build connection string
		fmt.Println(dbURI)

		conn, err := gorm.Open("postgres", dbURI)
		if err != nil {
			fmt.Print(err)
		}

		db = conn
		// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/film", film.GetFilms)
	app.Get("/api/v1/film/:id", film.GetFilm)
	app.Post("/api/v1/film", film.NewFilm)
	app.Delete("api/v1/film/:id", film.DelFilm)

}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen("127.0.0.1:8080")
}
