package cmd

import (
	"fmt"
	"log"

	"github.com/MirzaKarabulut/fill-labs/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateServer(port int) {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	routes.Router(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
