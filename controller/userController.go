package controller

import (
	"github.com/ThePratikSah/silent-flag/database"
	"github.com/ThePratikSah/silent-flag/helper"
	"github.com/ThePratikSah/silent-flag/model"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		return helper.SendResponse(c, 400, err.Error(), nil)
	}

	if err := db.Create(&user).Error; err != nil {
		return helper.SendResponse(c, 500, err.Error(), nil)
	}
	return helper.SendResponse(c, 200, "User Created Successfully", user)
}

func GetUsers(c *fiber.Ctx) error {
	return c.SendString("Get Users")
}

func GetUser(c *fiber.Ctx) error {

	return c.SendString("Get User")
}

func UpdateUser(c *fiber.Ctx) error {

	return c.SendString("Update User")
}

func DeleteUser(c *fiber.Ctx) error {

	return c.SendString("Delete User")
}
