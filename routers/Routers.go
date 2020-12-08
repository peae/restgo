package routers

import (
	"restgo/controller"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		crudUser := new(controller.CrudUser)
		v1.GET("/users", crudUser.Index)
		v1.GET("/user/:id", crudUser.FindUserById)
		v1.DELETE("/user/:id", crudUser.DeleteUserById)
		v1.POST("/user", crudUser.AddUser)
		v1.PATCH("/user", crudUser.UpdateUser)
		v1.PUT("/user", crudUser.UpdateUser)
	}
	return router

	//	router := gin.Default()
	//
	//	router.GET("/", Index)
	//	router.GET("/users", FindAllUsers)
	//	router.GET("/user/:id", FindUserById)
	//	router.POST("/user", AddUser)
	//	router.PUT("/user/:id", UpdateUserById)
	//	router.DELETE("/user/:id", DeleteUserById)
	//
	//	return router
}

