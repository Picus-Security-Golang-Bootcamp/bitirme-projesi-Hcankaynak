package main

import (
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/product"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/user"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/config"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/database"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/logger"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/pkg/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

func main() {

	// Setting up environment variables, reading from config
	cfg, err := config.LoadConfig("./pkg/config/config")
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	// initialize logger
	logger.NewLogger(&cfg.Logger)
	defer logger.Close()

	// Connecting to database.
	DB := database.Connect(&cfg.DBConfig)

	// init gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	initHandlers(r, &cfg.ServerConfig, DB)

	server.StartServer(&cfg.ServerConfig, r)
}

func toDos() {
	// + project template
	// + create models (product, user, category)
	// + gin and server added.
	// + add handler, + creating router group,  + add logger
	// add jwt and middleware
	// add bulk csv
	// add basic services
	// add basket services
	// add swagger
	// add uuid
	// add advanced readme (brief explanation about project structure will be seemed complex)
	// add tests
}

func initHandlers(r *gin.Engine, cfg *config.ServerConfig, db *gorm.DB) {
	// initializing routers.
	zap.L().Debug("initializing routers")

	rootRouter := r.Group(cfg.RoutePrefix)
	productRouter := rootRouter.Group("/products")
	userRouter := rootRouter.Group("/users")
	//categoryRouter := rootRouter.Group("/categories")
	//basketRouter := rootRouter.Group("/baskets")

	productRepo := product.NewProductRepository(db)
	productRepo.Migration()
	product.NewProductHandler(productRouter, productRepo)

	userRepo := user.NewUserRepository(db)
	userRepo.Migration()
	user.NewUserHandler(userRouter, userRepo)
}
