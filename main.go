package main

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	port = ":3000"
)

type Item struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Price float32 `json:"price"`
}

var items []Item = []Item{
	{Id: 0, Title: "First item", Price: 10.2},
	{Id: 1, Title: "Second item", Price: 5.1},
	{Id: 2, Title: "Third item", Price: 4.5},
	{Id: 3, Title: "Fourth item", Price: 8.0},
	{Id: 4, Title: "Fifth item", Price: 0.5},
}

func getItems(c *fiber.Ctx) error {
	fmt.Println(c.Method(), c.Path())
	return c.JSON(items)
}

func getSingleItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Incorrect id")
	}

	for _, item := range items {
		if item.Id == id {
			return c.JSON(item)
		}
	}

	return c.Status(http.StatusNotFound).SendString("Item not found")
}

func main() {
	app := fiber.New()

	app.Get("/items", getItems)
	app.Get("/items/:id", getSingleItem)

	err := app.Listen(port)
	if err != nil {
		fmt.Println("server start error", err)
	}
}