package routes

import (
	"angellisandroerazo/personal-blog/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html/v2"
)

func App() *fiber.App {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "/static/assets")

	app.Get("/", controllers.Index)

	app.Get("/post/:id", controllers.ViewPost)

	admin := app.Group("/admin")

	admin.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "admin",
		},
	}))

	admin.Get("/", controllers.AdminIndex)

	admin.Get("/create", controllers.CreatePage)
	admin.Post("/create-post", controllers.CreatePost)

	admin.Get("/edit/:id", controllers.UpdatePage)
	admin.Post("/edit-post/:id", controllers.UpdatePost)

	admin.Get("/delete/:id", controllers.DeletePost)

	return app
}
