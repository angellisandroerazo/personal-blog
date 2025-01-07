package main

import (
	"angellisandroerazo/personal-blog/routes"
	"log"
)

func main() {
	app := routes.App()

	log.Fatal(app.Listen(":3000"))
}
