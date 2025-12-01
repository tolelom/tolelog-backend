package main

import (
	"log"
	"tolelom_api/internal/config"
	"tolelom_api/internal/router"

	"github.com/gofiber/fiber/v2"

	_ "tolelom_api/docs"
)

func main() {
	cfg := config.LoadConfig()

	if err := config.InitDataBase(cfg); err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"status":  "error",
				"message": err.Error(),
				"data":    nil,
			})
		},
	})

	router.Setup(app)

	addr := ":" + cfg.Port
	log.Printf("Server listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatalf("Fiber server failed to start: %v", err)
	}
}
