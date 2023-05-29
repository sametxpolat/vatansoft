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

type ProductHandler struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) Stocks(c echo.Context) error {
	stocks, err := h.productService.Stocks()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "not found products")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, stocks)
}

func (h *ProductHandler) Stock(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	product, err := h.productService.Stock(uint(id))
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, "not found product")
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(c echo.Context) error {
	var product dto.CProduct

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong request")
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.productService.Create(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "create product")
}

func (h *ProductHandler) Update(c echo.Context) error {
	var product dto.UProduct

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "wrong request")
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.productService.Update(uint(id), &product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "update product")
}

func (h *ProductHandler) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	err = h.productService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found product")
	}

	return c.JSON(http.StatusOK, "delete product")
}

func (h *ProductHandler) Filter(c echo.Context) error {
	var req dto.Filter

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	products, err := h.productService.Filter(req.Name, req.Barcode, req.Price, req.CategoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, products)
}
