package handler

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"githuh.com/1995parham-teaching/students/internal/http/request"
	"githuh.com/1995parham-teaching/students/internal/model"
	"githuh.com/1995parham-teaching/students/internal/store/student"
	"go.uber.org/zap"
)

type Student struct {
	Store  student.Student
	Logger *zap.Logger
}

func (s *Student) List(c *fiber.Ctx) error {
	l, err := s.Store.GetAll()
	if err != nil {
		s.Logger.Error("store.getall failed", zap.Error(err))

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(l)
}

func (s *Student) Get(c *fiber.Ctx) error {
	idStr := c.Params("id", "-")
	if idStr == "-" {
		s.Logger.Info("cannot get id", zap.String("url", c.OriginalURL()))

		return fiber.ErrBadRequest
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		s.Logger.Info("cannot parse id", zap.String("id", idStr))

		return fiber.ErrBadRequest
	}

	std, err := s.Store.Get(id)
	if err != nil {
		if errors.Is(err, student.ErrStudentNotFound) {
			return fiber.ErrNotFound
		}

		s.Logger.Error("store.get", zap.Error(err), zap.Int64("id", id))

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(std)
}

func (s *Student) Create(c *fiber.Ctx) error {
	var req request.Student

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return fiber.ErrBadRequest
	}

	m := model.Student{
		ID:        rand.Int63(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Average:   0.0,
		Age:       req.Age,
		Courses:   []model.Course{},
	}

	if err := s.Store.Set(m); err != nil {
		log.Println(err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(m)
}

func (s *Student) Register(g fiber.Router) {
	g.Get("/", s.List)
	g.Post("/", s.Create)
	g.Get("/:id", s.Get)
}
