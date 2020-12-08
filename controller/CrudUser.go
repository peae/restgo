package controller

import (
	"restgo/database"
	"restgo/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CrudUser struct {
	Basic
}

func (u *CrudUser) Index(c *gin.Context)  {
	var users []models.User
	//database.DB.Select("id, name").Order("id").Find(&users)
	database.DB.Find(&users)
	fmt.Println(users)
	u.JsonSuccess(c, http.StatusOK, gin.H{"data":users})
}

func (u *CrudUser) FindUserById(c *gin.Context) {
	var user models.User

	if database.DB.First(&user, c.Param("id")).Error != nil {
		u.JsonFail(c, http.StatusNotFound, "用户不存在")
		return
	}
	u.JsonSuccess(c, http.StatusCreated, gin.H{"data": user})
}

func (u *CrudUser) DeleteUserById(c *gin.Context) {
	var user models.User

	if database.DB.First(&user, c.Param("id")).Error != nil {
		u.JsonFail(c, http.StatusNotFound, "用户不存在")
		return
	}
	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		u.JsonFail(c, http.StatusBadRequest, err.Error())
		return
	}

	u.JsonSuccess(c, http.StatusCreated, gin.H{"data": "删除成功"})
}

func (u *CrudUser) AddUser(c *gin.Context)  {

	var user models.User

	if err := c.ShouldBindJSON(&user); err == nil {
		fmt.Println(user)

		if database.DB.First(&user, user.Id).Error == nil {
			u.JsonFail(c, http.StatusNotFound, "用户已存在")
			return
		}

		if err := database.DB.Create(&user).Error; err != nil {
			u.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		u.JsonSuccess(c, http.StatusCreated, gin.H{"message": "创建成功"})
	}else {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		u.JsonFail(c, http.StatusBadRequest, "创建失败")
	}
}

type User2 struct {
	Id	string
	Name	string
}


func (u *CrudUser) UpdateUser(c *gin.Context) {

	var user models.User
	var user2 User2

	if err := c.ShouldBindJSON(&user2); err == nil {
		fmt.Println(user2)

		if database.DB.First(&user, user2.Id).Error != nil {
			u.JsonFail(c, http.StatusNotFound, "用户不存在")
			return
		}
		fmt.Println(user)

		user.Name = user2.Name

		fmt.Println(user)

		if err := database.DB.Save(&user).Error; err != nil {
			u.JsonFail(c, http.StatusBadRequest, err.Error())
			return
		}

		u.JsonSuccess(c, http.StatusCreated, gin.H{"message": "更新成功"})
	}else {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		u.JsonFail(c, http.StatusBadRequest, "更新失败")
	}


}








