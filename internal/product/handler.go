package product

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type productHandler struct {
	repo *ProductRepository
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepository) {
	zap.L().Debug("Product Handler initializing..")

	//h := &productHandler{repo: repo}

	zap.L().Debug("User Handler initialized")
}
