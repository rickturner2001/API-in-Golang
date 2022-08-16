package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rickturner2001/crm-test/controllers"
	"github.com/rickturner2001/crm-test/database"
	"github.com/rickturner2001/crm-test/models"
)

func SetRoutes(app *fiber.App){
	
	app.Get("/", controllers.Home)
	
	// Auth
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
	app.Get("/user", controllers.User)
	
	// API
	app.Get("/api/v1/lead", func(c *fiber.Ctx) error{
	db := database.DBConection
	var leads []models.Lead
	db.Find(&leads)
	return c.JSON(leads)})

	app.Get("/api/v1/lead/:id",func (c *fiber.Ctx) error {
	id := c.Params("id")
	db  := database.DBConection
	var lead models.Lead
	db.Find(&lead, id)
	return c.JSON(lead)

})
	app.Post("/api/v1/lead", func(c *fiber.Ctx) error {
	db := database.DBConection
	lead := new(models.Lead)
	if err := c.BodyParser(lead); err != nil{
		return err
	}
	db.Create(&lead)
	return c.JSON(lead)

})
	app.Delete("/api/v1/lead/:id" ,func (c *fiber.Ctx) error{

	id := c.Params("id")
	db := database.DBConection

	var lead models.Lead
	db.First(&lead, id)
	if lead.Name == ""{
		return db.Error
	}

	db.Delete(&lead)
	return c.JSON("Lead Successfully deleted")
	
})

}