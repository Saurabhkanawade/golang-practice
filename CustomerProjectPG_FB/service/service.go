package service

import (
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"github.com/saurabhkanawade/customerprojectpg_fb/database"
	"github.com/saurabhkanawade/customerprojectpg_fb/model"
	"log"
)

func GetCustomers(ctx *fiber.Ctx) {
	var customers []model.Customer
	log.Println("fetching the customers..........")

	db := database.Connect()
	defer db.Close()

	err := db.Model(&customers).Select()
	if err != nil {
		return
	}

	ctx.JSON(customers)
}

func GetCustomer(ctx *fiber.Ctx) {
	log.Println("fetching customer ...........")

	db := database.Connect()
	defer db.Close()

	id := ctx.Params("id")

	customer := &model.Customer{
		Id: id,
	}

	err := db.Model(customer).WherePK().Select()
	if err != nil {
		log.Fatalln(err)
	}

	ctx.JSON(customer)

	log.Println("fetch success with customer id :", customer.Id)
}

func CreateCustomer(ctx *fiber.Ctx) {
	log.Println("Creating new customer.........")

	db := database.Connect()
	defer db.Close()

	var customer = model.Customer{
		Id: uuid.New().String(),
	}

	err := ctx.BodyParser(&customer)
	if err != nil {
		return
	}

	_, err = db.Model(customer).Insert()
	if err != nil {
		return
	}

	err = ctx.JSON(customer)
	if err != nil {
		return
	}
}

func UpdateCustomer(ctx *fiber.Ctx) {

	log.Println("Updating the customer.........")

	db := database.Connect()
	defer db.Close()

	id := ctx.Params("id")

	customer := &model.Customer{
		Id: id,
	}

	err := ctx.BodyParser(&customer)
	if err != nil {
		return
	}

	_, err2 := db.Model(customer).WherePK().Update()
	if err2 != nil {
		log.Fatalln(err2)
	}

	err3 := ctx.JSON(customer)
	if err3 != nil {
		log.Fatalln(err3)
	}

}

func DeleteCustomer(ctx *fiber.Ctx) {
	log.Println("Deleting the customer.............")

	db := database.Connect()
	defer db.Close()

	//var customer model.Customer
	//id := ctx.Params("id")
	//
	//err := db.Model(&customer)

}
