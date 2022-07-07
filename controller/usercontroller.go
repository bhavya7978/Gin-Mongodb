package controller

import (
	"net/http"
	"project_mongodb-go/entity"
	"project_mongodb-go/service"

	"github.com/gin-gonic/gin"
)

//STRUCT
type UserController struct {
	UserService service.UserServices
}

//CONSTRUCTOR
func New(userservice service.UserServices) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetUser(c *gin.Context) {
	username := c.Param("name")
	user, err := uc.UserService.GetUser(&username)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAll(c *gin.Context) {
	users, err := uc.UserService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(c *gin.Context) {

	username := c.Param("name")
	err := uc.UserService.DeleteUser(&username)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userroute := rg.Group("/user")
	userroute.POST("/createuser", uc.CreateUser)
	userroute.GET("/getuser/:name", uc.GetUser)
	userroute.GET("/getalluser", uc.GetAll)
	userroute.PATCH("/updateuser", uc.UpdateUser)
	userroute.DELETE("/deleteuser/:name", uc.DeleteUser)

}
