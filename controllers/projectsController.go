package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rickturner2001/crm-test/database"
	"github.com/rickturner2001/crm-test/models"
)

func NewProject(c *fiber.Ctx) error{
	db := database.DBConection
	var data map[string]string
	if err:= c.BodyParser(&data); err != nil{
		return err
	}

	state, err := strconv.Atoi(data["state"]);
	if err != nil{
		return err
	} 

	userID, err := strconv.Atoi(data["appointee"]);
	if err != nil{
		return err
	}

	priority, err := strconv.Atoi(data["priority"]);
	if err != nil{
		return err
	}

	layout := "2006-01-02"
	due := data["due"]

	dueDate, err := time.Parse(layout, due)
	if err != nil{
		return err
	}

	currentTime := time.Now()
	currentTime.Format("2006-01-02")


	project := models.Project{
		Task: data["task"],
		State: state,
		UserID: userID,
		Priority: priority,
		Due: dueDate,
		Created: currentTime,
	}

	db.Create(&project)
	return c.JSON(project)

}


func GetProjects(c *fiber.Ctx) error{
	db := database.DBConection
	var projects []models.Project
	db.Find(&projects)
	return c.JSON(projects)
}


func PatchState(c *fiber.Ctx) error{
	db := database.DBConection
	var data map[string]string
	if err:= c.BodyParser(&data); err != nil{
		return err
	}
	id, err := strconv.Atoi(data["ID"]);
	if err != nil{
		return err
	}
	state, err := strconv.Atoi(data["state"]);
	if err != nil{
		return err
	}

	db.Model(&models.Project{}).Where("ID = ?", id).Update("state", state)
	return c.JSON(db.Model(&models.Project{}).Where("ID = ?", id))
}

func PatchPriority(c *fiber.Ctx) error{
	db := database.DBConection
	var data map[string]string
	if err:= c.BodyParser(&data); err != nil{
		return err
	}
	id, err := strconv.Atoi(data["ID"]);
	if err != nil{
		return err
	}
	priority, err := strconv.Atoi(data["priority"]);
	if err != nil{
		return err
	}

	db.Model(&models.Project{}).Where("ID = ?", id).Update("priority", priority)
	return c.JSON(db.Model(&models.Project{}).Where("ID = ?", id))
}