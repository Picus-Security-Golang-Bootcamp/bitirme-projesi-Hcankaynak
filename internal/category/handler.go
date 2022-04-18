package category

import (
	"fmt"
	pagination "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal"
	"github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/api"
	httpErr "github.com/Picus-Security-Golang-Bootcamp/bitirme-projesi-Hcankaynak/internal/httpErrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type categoryHandler struct {
	repo *CategoryRepository
}

func NewCategoryHandler(r *gin.RouterGroup, repo *CategoryRepository) {
	c := categoryHandler{
		repo: repo,
	}
	c.repo.Migration()

	r.GET("/list", c.listAllCategories)
	r.POST("/", c.createCategory)
}

// listAllCategories listing all categories with pagination
func (ca *categoryHandler) listAllCategories(c *gin.Context) {
	pageIndex, pageSize := pagination.GetPaginationParametersFromRequest(c)
	fmt.Println(pageIndex, pageSize)
	categories, count := ca.repo.GetAllCategories(pageIndex, pageSize)

	paginatedResult := pagination.NewFromGinRequest(c, count)
	paginatedResult.Items = &categories

	c.JSON(http.StatusOK, paginatedResult)
}

// createCategory create an individual category
func (ca *categoryHandler) createCategory(c *gin.Context) {
	var req api.Category
	if err := c.Bind(&req); err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusBadRequest, "check your request body", nil)))
		return
	}

	category, err := ca.repo.createCategory(&req)
	if err != nil {
		c.JSON(httpErr.ErrorResponse(httpErr.NewRestError(http.StatusInternalServerError, "category couldn't created", nil)))
		return
	}

	c.JSON(http.StatusOK, category)
}
