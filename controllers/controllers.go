package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	posts, message := listPosts()

	return c.Render("index", fiber.Map{
		"message": message,
		"Posts":   posts,
	})
}

func AdminIndex(c *fiber.Ctx) error {

	posts, message := listPosts()

	return c.Render("admin/index", fiber.Map{
		"message": message,
		"Posts":   posts,
	})
}
