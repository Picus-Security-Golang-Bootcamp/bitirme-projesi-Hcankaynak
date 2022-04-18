package product

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	httpErr "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productHandler struct {
	repo *ProductRepository
	cfg  *config.JWTConfig
}

func NewProductHandler(r *gin.RouterGroup, repo *ProductRepository, cfg *config.JWTConfig) {
	h := &productHandler{repo: repo, cfg: cfg}
	h.repo.Migration()

	r.Use(middleware.AuthMiddleware(cfg.SecretKey))
	r.POST("/", h.createProduct)

}

func (p *productHandler) createProduct(c *gin.Context) {
	// TODO: swagger product item.
	var req api.Product
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
		return
	}

	prod, err := p.repo.createProduct(*ResponseToProduct(&req))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "product couldn't added", err)))
		return
	}

	c.JSON(http.StatusOK, prod)
}
