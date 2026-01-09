package dish

import "github.com/gin-gonic/gin"

type DishHandler struct{}

func (h *DishHandler) List(c *gin.Context) {}

func (h *DishHandler) Create(c *gin.Context) {}

func (h *DishHandler) Delete(c *gin.Context) {}

func NewDishHandler() *DishHandler {
	return &DishHandler{}
}
