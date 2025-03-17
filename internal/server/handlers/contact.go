package handlers

import (
	"golang-microservices-course-spring-2025/internal/server/models"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetContact(ctx *fiber.Ctx) error {
    log.Println("Get Contact")
	birthdate, _ := time.Parse("02-01-2006", "10-10-2006")
	return ctx.JSON(models.Contact{
		ID:         1,
		Username:   "Pety",
		GivenName:  "Cheater",
		FamilyName: "Moskovets",
		Phones:     []models.Phone{},
		Email:      []string{"petyMoscow@ya.ru"},
		Birthdate:  birthdate,
	})
}

func CreateContact(ctx *fiber.Ctx) error {
    log.Println("Create Contact")
	var contact models.Contact
	return ctx.BodyParser(&contact)
}

func UpdateContact(ctx *fiber.Ctx) error {
    log.Println("Update Contact")
	var contact models.Contact
	return ctx.BodyParser(&contact)
}

func DeleteContact(ctx *fiber.Ctx) error {
    log.Println("Delete Contact")
	return ctx.SendString("Contact was successfully deleted!")
}
