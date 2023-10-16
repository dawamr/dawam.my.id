package bootstrap

import (
	"github.com/dawamr/resume-cv-2023/pkg/database"
	"github.com/dawamr/resume-cv-2023/pkg/env"
	"github.com/dawamr/resume-cv-2023/pkg/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func NewApplication() *fiber.App {
	env.SetupEnvFile()
	database.SetupDatabase()
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:        engine,
		AppName:      "Resume Dawam Raja",
		ServerHeader: "Fiber",
	})
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/dashboard", monitor.New())
	router.InstallRouter(app)

	return app
}
