package user

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	httpErr "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/httpErrors"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	jwt_service "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
	"time"
)

type userHandler struct {
	repo *UserRepository
	cfg  config.JWTConfig
}

func NewUserHandler(r *gin.RouterGroup, repo *UserRepository, cfg config.JWTConfig) {

	h := userHandler{
		cfg:  cfg,
		repo: repo,
	}
	r.POST("/signUp", h.SignUp)
	r.POST("/login", h.Login)

}

func (h userHandler) SignUp(c *gin.Context) {
	var req api.User
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
		return
	}

	// looking for user with same email.
	user, _ := h.repo.getUserByEmail(req.Email)

	if user != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "user already exist", nil)))
		return
	}

	// creating new user
	newUser, err := h.repo.Create(&req)
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "new user cannot created", nil)))
		return
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": newUser.Email,
		"iat":   time.Now().Unix(),
		"iss":   os.Getenv("ENV"),
		"exp":   time.Now().Add(time.Duration(h.cfg.SessionTime) * time.Hour).Unix(),
		"roles": newUser.Role,
	})
	token := jwt_service.GenerateToken(jwtClaims, h.cfg.SecretKey)
	c.JSON(http.StatusOK, token)
}

func (h userHandler) Login(c *gin.Context) {
	var req api.User
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
		return
	}

	user, err := h.repo.Login(&req)
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "user don't exist", nil)))
		return
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iat":   time.Now().Unix(),
		"iss":   os.Getenv("ENV"),
		"exp":   time.Now().Add(time.Duration(h.cfg.SessionTime) * time.Hour).Unix(),
		"roles": user.Role.toString(),
	})

	token := jwt_service.GenerateToken(jwtClaims, h.cfg.SecretKey)
	c.JSON(http.StatusOK, token)

}
