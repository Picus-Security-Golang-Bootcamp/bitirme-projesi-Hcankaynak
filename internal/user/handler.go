package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userHandler struct {
	repo *UserRepository
}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepository) {
	zap.L().Debug("User Handler initializing..")
}
