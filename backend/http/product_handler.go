package http

import (
	"errors"
	"glow-fx/domain"
	"glow-fx/response"
	"glow-fx/usecase"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecase.ProductUseCase
}

func NewProductHandler(u *usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	search := c.Query("search")
	status := c.Query("status")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	data, total, err := h.usecase.GetProducts(search, status, page, limit)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, gin.H{
		"data":       data,
		"total":      total,
		"page":       page,
		"limit":      limit,
		"total_page": int(math.Ceil(float64(total) / float64(limit))),
	})
}

func (h *ProductHandler) GetProductById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(errors.New("invalid id"))
		return
	}

	data, err := h.usecase.GetProductById(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, data)
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var input domain.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	data, err := h.usecase.CreateProduct(input)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, data)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(errors.New("invalid id"))
		return
	}

	var input domain.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	data, err := h.usecase.UpdateProduct(uint(id), input)
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, data)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Error(errors.New("invalid id"))
		return
	}

	err = h.usecase.DeleteProduct(uint(id))
	if err != nil {
		c.Error(err)
		return
	}

	response.Success(c, gin.H{
		"message": "product deleted successfully",
	})
}
