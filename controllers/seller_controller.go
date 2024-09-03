package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new-shout-golang/models"
	"new-shout-golang/services"
	"new-shout-golang/utils/resp"
	"strconv"
)

type SellerController struct {
	Service services.SellerService
}

func NewSellerController(service services.SellerService) *SellerController {
	return &SellerController{Service: service}
}

func (c *SellerController) CreateSeller(ctx *gin.Context) {
	var seller models.Seller
	if err := ctx.BindJSON(&seller); err != nil {
		resp.Error(ctx, http.StatusBadRequest, "Invalid input")
		return
	}

	createdSeller, err := c.Service.CreateSeller(&seller)
	if err != nil {
		resp.Error(ctx, http.StatusInternalServerError, "Failed to create seller")
		return
	}

	resp.Success(ctx, createdSeller)
}

func (c *SellerController) GetAllSellers(ctx *gin.Context) {
	sellers, err := c.Service.GetAllSellers()
	if err != nil {
		resp.Error(ctx, http.StatusInternalServerError, "Failed to get sellers")
		return
	}

	resp.Success(ctx, sellers)
}

func (c *SellerController) GetSellerByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		resp.Error(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}

	seller, err := c.Service.GetSellerByID(uint(id))
	if err != nil {
		resp.Error(ctx, http.StatusNotFound, "Seller not found")
		return
	}

	resp.Success(ctx, seller)
}

func (c *SellerController) UpdateSeller(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		resp.Error(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}

	var seller models.Seller
	if err := ctx.BindJSON(&seller); err != nil {
		resp.Error(ctx, http.StatusBadRequest, "Invalid input")
		return
	}
	seller.ID = uint(id)

	updatedSeller, err := c.Service.UpdateSeller(&seller)
	if err != nil {
		resp.Error(ctx, http.StatusInternalServerError, "Failed to update seller")
		return
	}

	resp.Success(ctx, updatedSeller)
}

func (c *SellerController) DeleteSeller(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		resp.Error(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}

	if err := c.Service.DeleteSeller(uint(id)); err != nil {
		resp.Error(ctx, http.StatusNotFound, "Seller not found")
		return
	}

	resp.Success(ctx, nil)
}
