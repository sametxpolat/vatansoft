package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sametxpolat/vatansoft/dto"
	"github.com/sametxpolat/vatansoft/service"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Categories(c echo.Context) error {
	categories, err := h.categoryService.Categories()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "not found categories")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) Category(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	category, err := h.categoryService.Category(uint(id))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "not found category")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, category)
}

func (h *CategoryHandler) Create(c echo.Context) error {
	var category dto.CCategory

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong request")
	}

	err = json.Unmarshal(body, &category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.categoryService.Create(&category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "create category")
}

func (h *CategoryHandler) Update(c echo.Context) error {
	var category dto.UCategory

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong request")
	}

	err = json.Unmarshal(body, &category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.categoryService.Update(uint(id), &category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "update category")
}

func (h *CategoryHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.categoryService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found category")
	}

	return c.JSON(http.StatusOK, "delete category")
}
