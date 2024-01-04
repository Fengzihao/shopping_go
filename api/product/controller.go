package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopping_go/domain/product"
	"shopping_go/utils/api_helper"
	"shopping_go/utils/pagination"
)

type Controller struct {
	productService product.Service
}

func NewProductController(server product.Service) *Controller {
	return &Controller{
		productService: server,
	}
}

// GetProducts
// @Summary 获得商品列表（分页）
// @Tags Product
// @Accept json
// @Produce json
// @Param qt query string false "Search text to find matched sku numbers and names"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /product [get]
func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	query := g.Query("qt")
	if query != "" {
		page = c.productService.SearchProduct(query, page)
	} else {
		page = c.productService.GetAll(page)
	}
	g.JSON(http.StatusOK, page)
}

// CreateProduct
// @Summary 创建商品
// @Tags Product
// @Accept json
// @Produce json
// @Param Authorization  header    string  true  "Authentication header"
// @Param CreateProductRequest body CreateProductRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /product [post]
func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err := c.productService.CreateProduct(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{
		Message: "Product Created",
	})

}

// DeleteProduct
// @Summary 删除商品根据sku
// @Tags Product
// @Accept json
// @Produce json
// @Param DeleteProductRequest body DeleteProductRequest true "sku of product"
// @Param Authorization header    string  true  "Authentication header"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /product [delete]
func (c *Controller) DeleteProduct(g *gin.Context) {
	var req DeleteProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err := c.productService.DeleteProduct(req.SKU)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Message: "Product Deleted"})
}

// UpdateProduct
// @Summary 更新商品更加sku
// @Tags Product
// @Accept json
// @Produce json
// @Param UpdateProductRequest body UpdateProductRequest true "product information"
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {object} CreateProductResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /product [patch]
func (c *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	err := c.productService.UpdateProduct(req.ToProduct())
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusOK, api_helper.Response{Message: "Product Updated"})
}
