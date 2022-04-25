package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"githuh.com/cng-by-example/students/internal/http/handler"
	"githuh.com/cng-by-example/students/internal/store/student"
)

func main() {
	app := fiber.New()

	hnd := handler.Student{
		Store: student.NewInMemory(),
	}

	hnd.Register(app.Group("/students"))

	if err := app.Listen(":1373"); err != nil {
		log.Fatal(err)
	}
}
