package handlers

import (
	"golang-microservices-course-spring-2025/internal/server/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetGroup(ctx *fiber.Ctx) error {
    log.Println("Get Group")
	return ctx.JSON(models.Group{
		ID:          1,
		Title:       "test group",
		Description: "group for test functionality",
		Contacts:    []int{},
	})
}

func CreateGroup(ctx *fiber.Ctx) error {
    log.Println("Create Group")
	var group models.Group
	return ctx.BodyParser(&group)
}

func UpdateGroup(ctx *fiber.Ctx) error {
    log.Println("Update Group")
	var group models.Group
	return ctx.BodyParser(&group)
}

func DeleteGroup(ctx *fiber.Ctx) error {
    log.Println("Delete Group")
	return ctx.SendString("Group was successfully deleted!")
}
