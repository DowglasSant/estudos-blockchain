package handler

import "github.com/gin-gonic/gin"

func NewRouter(h *BlockchainHandler, webDir string) *gin.Engine {
	r := gin.Default()

	r.StaticFile("/", webDir+"/index.html")

	r.GET("/blocks", h.getBlocks)
	r.POST("/blocks", h.mineBlock)
	r.GET("/chain/valid", h.validateChain)

	return r
}
