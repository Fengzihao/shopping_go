package user

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"shopping_go/config"
	"shopping_go/domain/user"
	"shopping_go/utils/api_helper"
	jwtHelper "shopping_go/utils/jwt"
	"strconv"
	"time"
)

type Controller struct {
	userService *user.Service
	appConfig   *config.Configuration
}

func NewUserController(service *user.Service, appConfig *config.Configuration) *Controller {
	return &Controller{
		userService: service,
		appConfig:   appConfig,
	}
}

// CreateUser
// @Summary 创建用户
// @Tags Auth
// @Accept json
// @Produce json
// @Param CreateUserRequest body CreateUserRequest true "user information"
// @Success 201 {object} CreateUserResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user [post]
func (c *Controller) CreateUser(g *gin.Context) {
	var req CreateUserRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	newUser := user.NewUser(req.Username, req.Password, req.Password2)
	err := c.userService.Create(newUser)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(http.StatusCreated, CreateUserResponse{
		Username: newUser.Username,
	})
}

// VerifyToken 验证token
func (c *Controller) VerifyToken(g *gin.Context) {
	token := g.GetHeader("Authorization")
	verifyToken := jwtHelper.VerifyToken(token, c.appConfig.SecretKey)
	g.JSON(http.StatusOK, verifyToken)
}

// Login
// @Summary 登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "user information"
// @Success 200 {object} LoginResponse
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /user/login [post]
func (c *Controller) Login(g *gin.Context) {
	var req LoginRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	user, err := c.userService.GetUser(req.Username, req.Password)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	token := jwtHelper.VerifyToken(user.Token, c.appConfig.SecretKey)
	if token == nil {
		jwtToken := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(user.ID), 10),
				"username": user.Username,
				"iat":      time.Now().Unix(),
				"iss":      os.Getenv("ENV"),
				"exp":      time.Now().Add(24 * time.Hour).Unix(),
				"isAdmin":  user.IsAdmin,
			})
		tokenStr := jwtHelper.GenerateToken(jwtToken, c.appConfig.SecretKey)
		user.Token = tokenStr
		err := c.userService.UpdateUser(&user)
		if err != nil {
			api_helper.HandleError(g, err)
			return
		}
	}
	g.JSON(http.StatusOK, LoginResponse{
		Username: user.Username,
		UserId:   user.ID,
		Token:    user.Token,
	})
}
