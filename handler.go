package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// Handler
type StudentHandler struct {
	repo *StudentRepository
}

func NewStudentHandler(repo *StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

func (h *StudentHandler) GetStudents(c echo.Context) error {
	students, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, students)
}

func (h *StudentHandler) GetStudent(c echo.Context) error {
	id := c.Param("id")
	student, err := h.repo.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Student not found",
		})
	}
	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) CreateStudent(c echo.Context) error {
	var req StudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	student := Student{
		ID:          uuid.New(),
		Name:        req.Nama,
		Email:       req.SuratElektronik,
		PhoneNumber: req.NoHP,
		Address:     req.Alamat,
		GPA:         req.IPK,
		IsGraduate:  req.IsGraduate,
	}

	if err := h.repo.Create(student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, student)
}

func (h *StudentHandler) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	var req StudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	student := Student{
		Name:        req.Nama,
		Email:       req.SuratElektronik,
		PhoneNumber: req.NoHP,
		Address:     req.Alamat,
		GPA:         req.IPK,
		IsGraduate:  req.IsGraduate,
	}

	if err := h.repo.Update(id, student); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, student)
}

func (h *StudentHandler) DeleteStudent(c echo.Context) error {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Student deleted successfully",
	})
}