package router

import (
	"log"
	"net/http"

	"github.com/dongil91/module-test/domain"
	"github.com/gin-gonic/gin"
)

type SellerRouter struct {
	sellerService domain.SellerService
}

type CreateSellerRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type DeleteSellerRequest struct {
	Identity int `json:"identity" binding:"required"`
}

type UpdateSellerRequest struct {
	Identity int    `json:"identity" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

func NewSellerRouter(engine *gin.Engine, userService domain.SellerService) {
	sellerRouter := SellerRouter{sellerService: userService}

	engine.GET("/apis/sellers/me", sellerRouter.FindById)
	engine.GET("/apis/sellers", sellerRouter.FindAll)
	engine.POST("/apis/sellers", sellerRouter.Create)
	engine.PUT("/apis/sellers", sellerRouter.Update)
	engine.DELETE("apis/sellers", sellerRouter.Delete)
}

func (ur *SellerRouter) FindById(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")
	log.Print(authorization)
	me, _ := ur.sellerService.FindById(1)
	c.JSON(http.StatusOK, me)
}

func (ur *SellerRouter) FindAll(c *gin.Context) {
	users, _ := ur.sellerService.FindAll()
	c.JSON(http.StatusOK, users)
}

func (ur *SellerRouter) Create(c *gin.Context) {
	createUserRequest := &CreateSellerRequest{}
	err := c.ShouldBind(createUserRequest)
	res := make(map[string]string)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}
	log.Println(createUserRequest)
	err = ur.sellerService.Create(createUserRequest.Name, createUserRequest.Email)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["message"] = "success for creating new user"
	c.JSON(http.StatusCreated, res)
}

func (ur *SellerRouter) Delete(c *gin.Context) {
	deleteUserRequest := &DeleteSellerRequest{}
	err := c.ShouldBind(deleteUserRequest)
	res := make(map[string]string)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	log.Println(deleteUserRequest)
	err = ur.sellerService.Delete(deleteUserRequest.Identity)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["message"] = "success for delete user"
	c.JSON(http.StatusOK, res)
}

func (ur *SellerRouter) Update(c *gin.Context) {
	updateUserRequest := &UpdateSellerRequest{}
	err := c.ShouldBind(updateUserRequest)
	res := make(map[string]string)
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	log.Println(updateUserRequest)
	err = ur.sellerService.Update(updateUserRequest.Name, int64(updateUserRequest.Identity))
	if err != nil {
		res["error"] = err.Error()
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res["message"] = "success for update user"
	c.JSON(http.StatusOK, res)
}
