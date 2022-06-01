package controllers

import (
	"github.com/Anirudh-rao/RESTAPI-GOFIBER/models"
	"github.com/Kamva/mgm/v2"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
)

// GetAllTodos - GET /api/todos
// GetAllTodos - GET /api/todos
func GetAllTodos(ctx *fiber.Ctx) {
	collection := mgm.Coll(&models.Todo{})
	todos := []models.Todo{}

	err := collection.SimpleFind(&todos, bson.D{})
	if err != nil {
		ctx.Status(500).JSON(fiber.Map{
			"ok":    false,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":    true,
		"todos": todos,
	})
}

// GetTodoByID - GET /api/todos/:id
func GetTodoByID(ctx *fiber.Ctx) {
	id := ctx.Params("id")

	todo := &models.Todo{}
	collection := mgm.Coll(todo)

	err := collection.FindByID(id, todo)
	if err != nil {
		ctx.Status(404).JSON(fiber.Map{
			"ok":    false,
			"error": "Todo not found.",
		})
		return
	}

	ctx.JSON(fiber.Map{
		"ok":   true,
		"todo": todo,
	})
}
