package api

import (
	"log"
	cartApi "shopping_go/api/cart"
	categoryApi "shopping_go/api/category"
	orderApi "shopping_go/api/order"
	productApi "shopping_go/api/product"
	userApi "shopping_go/api/user"
	"shopping_go/config"
	"shopping_go/utils/middleware"

	"shopping_go/domain/cart"
	"shopping_go/domain/order"
	"shopping_go/domain/product"

	"shopping_go/domain/category"
	"shopping_go/domain/user"
	"shopping_go/utils/database_handler"

	"github.com/gin-gonic/gin"
)

// Databases 结构体
type Databases struct {
	categoryRepository    *category.Repository
	userRepository        *user.Repository
	productRepository     *product.Repository
	cartRepository        *cart.Repository
	cartItemRepository    *cart.ItemRepository
	orderRepository       *order.Repository
	orderedItemRepository *order.OrderedItemRepository
}

// 配置文件全局对象
var AppConfig = &config.Configuration{}

// 根据配置文件创建数据库
func CreateDBs() *Databases {
	cfgFile := "./config/config.yaml"
	conf, err := config.GetAllConfigValues(cfgFile)
	AppConfig = conf
	if err != nil {
		return nil
	}
	if err != nil {
		log.Fatalf("读取配置文件失败. %v", err.Error())
	}
	db := database_handler.NewMySQLDB(AppConfig.DatabaseSettings.DatabaseURI)
	return &Databases{
		categoryRepository:    category.NewCategoryRepository(db),
		cartRepository:        cart.NewCartRepository(db),
		userRepository:        user.NewUserRepository(db),
		productRepository:     product.NewProductRepository(db),
		cartItemRepository:    cart.NewCartItemRepository(db),
		orderRepository:       order.NewOrderRepository(db),
		orderedItemRepository: order.NewOrderedItemRepository(db),
	}
}

// 注册所有控制器
func RegisterHandlers(r *gin.Engine) {

	dbs := *CreateDBs()
	RegisterUserHandlers(r, dbs)
	RegisterCategoryHandlers(r, dbs)
	RegisterCartHandlers(r, dbs)
	RegisterProductHandlers(r, dbs)
	RegisterOrderHandlers(r, dbs)
}

// 注册分类控制器
func RegisterCategoryHandlers(r *gin.Engine, dbs Databases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := r.Group("/category")

	categoryGroup.POST(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), categoryController.CreateCategory)
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.POST(
		"/upload", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey),
		categoryController.BulkCreateCategory)
}

// 注册用户控制器
func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewUserService(*dbs.userRepository)
	userController := userApi.NewUserController(userService, AppConfig)
	userGroup := r.Group("/user")
	userGroup.POST("", userController.CreateUser)
	userGroup.POST("/login", userController.Login)

}

// 注册购物车控制器
func RegisterCartHandlers(r *gin.Engine, dbs Databases) {
	cartService := cart.NewService(*dbs.cartRepository, *dbs.cartItemRepository, *dbs.productRepository)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := r.Group("/cart", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.PATCH("/item", cartController.UpdateItem)
	cartGroup.GET("/", cartController.GetCart)
}

// 注册商品控制器
func RegisterProductHandlers(r *gin.Engine, dbs Databases) {
	productService := product.NewService(*dbs.productRepository)
	productController := productApi.NewProductController(*productService)
	productGroup := r.Group("/product")
	productGroup.GET("", productController.GetProducts)
	productGroup.POST(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.CreateProduct)
	productGroup.DELETE(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.DeleteProduct)
	productGroup.PATCH(
		"", middleware.AuthAdminMiddleware(AppConfig.JwtSettings.SecretKey), productController.UpdateProduct)

}

// 注册订单控制器
func RegisterOrderHandlers(r *gin.Engine, dbs Databases) {
	orderService := order.NewService(
		*dbs.orderRepository, *dbs.orderedItemRepository, *dbs.productRepository, *dbs.cartRepository,
		*dbs.cartItemRepository)
	productController := orderApi.NewOrderController(orderService)
	orderGroup := r.Group("/order", middleware.AuthUserMiddleware(AppConfig.JwtSettings.SecretKey))
	orderGroup.POST("", productController.CompleteOrder)
	orderGroup.DELETE("", productController.CancelOrder)
	orderGroup.GET("", productController.GetOrders)
}
