package recipe

import "github.com/gin-gonic/gin"

type RecipeHandler struct{}

func (h *RecipeHandler) ListByDish(c *gin.Context) {}

func (h *RecipeHandler) Create(c *gin.Context) {}

func (h *RecipeHandler) Get(c *gin.Context) {}

func (h *RecipeHandler) Update(c *gin.Context) {}

func (h *RecipeHandler) Delete(c *gin.Context) {}

func NewRecipeHandler() *RecipeHandler {
	return &RecipeHandler{}
}
