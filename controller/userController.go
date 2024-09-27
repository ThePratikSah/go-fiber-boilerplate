package controller

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {

	return c.SendString("Create User")
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
