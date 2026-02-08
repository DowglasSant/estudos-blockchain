package handler

import (
	"blockchain/blockchain/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlockchainHandler struct {
	svc *service.BlockchainService
}

func NewBlockchainHandler(svc *service.BlockchainService) *BlockchainHandler {
	return &BlockchainHandler{svc: svc}
}

func (h *BlockchainHandler) getBlocks(c *gin.Context) {
	c.JSON(http.StatusOK, h.svc.GetChain())
}

type mineBlockRequest struct {
	Data string `json:"data" binding:"required"`
}

func (h *BlockchainHandler) mineBlock(c *gin.Context) {
	var req mineBlockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	block := h.svc.MineBlock(req.Data)
	c.JSON(http.StatusCreated, block)
}

func (h *BlockchainHandler) validateChain(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"valid": h.svc.IsValid()})
}
