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

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type DeleteUserRequest struct {
	Identity int `json:"identity" binding:"required"`
}

type UpdateUserRequest struct {
	Identity int    `json:"identity" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func NewUserRouter(engine *gin.Engine, userService domain.UserService) {
	userRouter := UserRouter{userService: userService}

	engine.GET("/apis/users/me", userRouter.FindById)
	engine.GET("/apis/users", userRouter.FindAll)
	engine.POST("/apis/users", userRouter.Create)
	engine.PUT("/apis/users", userRouter.Update)
	engine.DELETE("apis/users", userRouter.Delete)
}

func (ur *UserRouter) FindById(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	log.Print(authorization)
	me, _ := ur.userService.FindById(1)
	c.JSON(http.StatusOK, me)
}

func (ur *UserRouter) FindAll(c *gin.Context) {
	users, _ := ur.userService.FindAll()
	c.JSON(http.StatusOK, users)
}

func (ur *UserRouter) Create(c *gin.Context) {
	createUserRequest := &CreateUserRequest{}
	err := c.ShouldBind(createUserRequest)
	res := make(map[string]string)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	log.Println(createUserRequest)
	err = ur.userService.Create(createUserRequest.Name, createUserRequest.Email)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["message"] = "success for creating new user"
	c.JSON(http.StatusCreated, res)
}

func (ur *UserRouter) Delete(c *gin.Context) {
	deleteUserRequest := &DeleteUserRequest{}
	err := c.ShouldBind(deleteUserRequest)
	res := make(map[string]string)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	log.Println(deleteUserRequest)
	err = ur.userService.Delete(deleteUserRequest.Identity)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["message"] = "success for delete user"
	c.JSON(http.StatusOK, res)
}

func (ur *UserRouter) Update(c *gin.Context) {
	updateUserRequest := &UpdateUserRequest{}
	err := c.ShouldBind(updateUserRequest)
	res := make(map[string]string)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	log.Println(updateUserRequest)
	err = ur.userService.Update(updateUserRequest.Name, int64(updateUserRequest.Identity))
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["message"] = "success for update user"
	c.JSON(http.StatusOK, res)
}
