package cart

import (
	"net/http"
	"shopping_go/domain/cart"
	"shopping_go/utils/api_helper"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	cartService *cart.Service
}

// NewCartController 初始化
func NewCartController(service *cart.Service) *Controller {
	return &Controller{
		cartService: service,
	}
}

// AddItem
// @Summary 添加Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [post]
func (c *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userID := api_helper.GetUserId(g)
	err := c.cartService.AddItem(userID, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Message: "item added to cart"})
}

// UpdateItem
// @Summary 更新Item
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart/item [patch]
func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userID := api_helper.GetUserId(g)
	err := c.cartService.UpdateItem(userID, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Message: "item Updated"})

}

// GetCart
// @Summary 获得购物车商品列表
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {array} ItemCartResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [get]
func (c *Controller) GetCart(g *gin.Context) {
	userID := api_helper.GetUserId(g)
	items, err := c.cartService.GetCartItems(userID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, ItemCartResponse{items: items})
}
