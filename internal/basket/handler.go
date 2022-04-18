package basket

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	httpErr "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/product"
	user2 "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	jwt_service "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/jwt"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type basketHandler struct {
	repo        *BasketRepository
	userRepo    *user2.UserRepository
	cfg         config.JWTConfig
	productRepo *product.ProductRepository
}

func NewBasketHandler(r *gin.RouterGroup, db *gorm.DB, cfg *config.JWTConfig, userRepo *user2.UserRepository, productRepo *product.ProductRepository) {
	h := &basketHandler{repo: NewBasketRepository(db), cfg: *cfg, userRepo: userRepo, productRepo: productRepo}
	h.repo.Migration()

	r.Use(middleware.AuthMiddleware(cfg.SecretKey))
	r.GET("/", h.getBasket)
	r.POST("/", h.addToBasket)
}

// getBasket listing all item in the basket.
func (b *basketHandler) getBasket(c *gin.Context) {
	token := jwt_service.VerifyToken(c.GetHeader(jwt_service.Authorization), b.cfg.SecretKey)

	products, err := b.repo.getBasketByUserId(token.Id)
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "no basket found", nil)))
		return
	}

	c.JSON(http.StatusOK, products)
}

// addToBasket first find item then add to basket
func (b *basketHandler) addToBasket(c *gin.Context) {
	token := jwt_service.VerifyToken(c.GetHeader(jwt_service.Authorization), b.cfg.SecretKey)

	var req api.AddToBasket
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
		return
	}

	productItem, err := b.productRepo.FindById(int(req.ID))
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "no product found", nil)))
		return
	}
	basket := Basket{Product: *productItem, UserID: token.Id, ProductID: int(productItem.ID)}

	err = b.repo.createBasketItem(basket)
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "no product found", err)))
		return
	}

	c.JSON(http.StatusOK, productItem)
}
