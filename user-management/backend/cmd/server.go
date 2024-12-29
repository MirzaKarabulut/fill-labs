package cmd

import (
	"fmt"
	"log"

	"github.com/MirzaKarabulut/fill-labs/internal/routes"
	"github.com/gofiber/fiber"
)

func CreateServer(port int) {
	app := fiber.New()

	routes.Router(app)

	log.Fatal(app.Listen(fmt.Sprintf("%d", port)))
}
