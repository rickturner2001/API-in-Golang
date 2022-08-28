package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rickturner2001/crm-test/database"
	"github.com/rickturner2001/crm-test/models"
	"github.com/rickturner2001/crm-test/routes"
)



func initDatabase(){
	var err error
	
	database.DBConection, err = gorm.Open("sqlite3", "data.db")
	if err != nil{
		panic("Failed to connect to database!")
	}
	fmt.Println("Successfully connected to database")
	database.DBConection.AutoMigrate(&models.Lead{})
	database.DBConection.AutoMigrate(&models.User{})
	database.DBConection.AutoMigrate(&models.Project{})
	fmt.Println("Database Migrated")
}


func main(){
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	initDatabase()
	routes.SetRoutes(app)
	app.Listen(":8090")
	defer database.DBConection.Close()

}
