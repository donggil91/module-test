package router

import (
	"log"
	"net/http"

	"github.com/dongil91/module-test/domain"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	userService domain.UserService
}

func NewUserRouter(engine *gin.Engine, userService domain.UserService) {
	userRouter := UserRouter{userService: userService}

	engine.GET("/apis/users/me", userRouter.FindById)
	engine.GET("/apis/users", userRouter.FindAll)
}

func (ur *UserRouter) FindById(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	log.Print(authorization)
	me, _ := ur.userService.FindById(c, 1)
	c.JSON(http.StatusOK, me)
}

func (ur *UserRouter) FindAll(c *gin.Context) {
	users, _ := ur.userService.FindAll()
	c.JSON(http.StatusOK, users)
}
