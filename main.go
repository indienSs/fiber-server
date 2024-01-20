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

func removeIndex[T any](s []T, index int) []T {
    return append(s[:index], s[index+1:]...)
}


func getItems(c *fiber.Ctx) error {
	fmt.Println(c.Method(), c.Path())
	return c.JSON(items)
}

func createItem(c *fiber.Ctx) error {
	var item Item
	
	if err := c.BodyParser(&item); err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid JSON")
	}

	item.Id = len(items)
	items = append(items, item)

	return c.JSON(item)
}

func getItem(c *fiber.Ctx) error {
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

func deleteItem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Incorrect id")
	}

	var deletedItem Item

	for i, item := range items {
		if item.Id == id {
			deletedItem = items[i]
			items = removeIndex(items, i)
		}
	}

	if &deletedItem == nil {
		return c.Status(http.StatusNotFound).SendString("Element not found")
	}

	return c.JSON(deletedItem)
}

func main() {
	app := fiber.New()

	app.Get("/items", getItems)
	app.Post("/items", createItem)
	app.Get("/items/:id", getItem)
	app.Delete("/items/:id", deleteItem)

	err := app.Listen(port)
	if err != nil {
		fmt.Println("server start error", err)
	}
}