package basket

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type basketHandler struct {
	repo *BasketRepository
}

func NewBasketHandler(r *gin.RouterGroup, repo *BasketRepository) {
	h := &basketHandler{repo: repo}

	r.GET("/", h.getAll)
}

func (b *basketHandler) getAll(c *gin.Context) {
	book := b.repo.getAll()

	c.JSON(http.StatusOK, book)
}
