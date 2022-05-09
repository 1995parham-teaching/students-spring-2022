package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/config"
	"githuh.com/cng-by-example/students/internal/http/handler"
	"githuh.com/cng-by-example/students/internal/store/student"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	hnd := handler.Student{
		Store:  student.NewInMemory(),
		Logger: logger.Named("handler.student"),
	}

	hnd.Register(app.Group("/students"))

	if err := app.Listen(cfg.Listen); err != nil {
		logger.Fatal("cannot listen", zap.String("listen", cfg.Listen),
			zap.Error(err))
	}
}
