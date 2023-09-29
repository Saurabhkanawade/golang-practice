package services

import (
	"github.com/gofiber/fiber"
	"github.com/mohdaalam005/fiber/database"
	"github.com/mohdaalam005/fiber/helpers"
	"github.com/mohdaalam005/fiber/models"
)

func GetStudents(c *fiber.Ctx) {
	var student []models.Student
	database.DB.Find(&student)
	c.JSON(student)
}

func GetStudent(c *fiber.Ctx) {
	id := c.Params("id")
	var student models.Student
	database.DB.Find(&student, id)
	c.JSON(student)
}

func CreateStudent(c *fiber.Ctx) {
	student := new(models.Student)
	err := c.BodyParser(student)
	helpers.CheckNilError(err)
	database.DB.Create(&student)
	c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) {
	id := c.Params("id")
	var student models.Student
	database.DB.Delete(&student, id)
	c.Send("successfully deleted !")
}
func UpdateStudent(c *fiber.Ctx) {
	id := c.Params("id")
	var student models.Student
	database.DB.Find(&student, id)
	c.BodyParser(&student)
	database.DB.Save(&student)
	c.JSON(student)
}
