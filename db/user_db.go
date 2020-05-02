package db

import (
	"aid.dev/goody/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var MysqlDB *gorm.DB

func InitDb() {
	mysqlDb, err := gorm.Open("mysql", "root:mysql_Jgt.2019@tcp(deva.wiki:3306)/goody?charset=utf8")

	if err != nil {
		log.Fatalln("failed to connect database: ", err)
	} else {
		log.Println("connect database success")
		mysqlDb.SingularTable(true)
		mysqlDb.AutoMigrate(&entity.UserInfo{})
		log.Println("create table success")
	}
	MysqlDB = mysqlDb
}

func CreateUser(c *gin.Context) {
	var user entity.UserInfo
	c.BindJSON(&user)
	fmt.Println(user)
	if err := MysqlDB.Create(&user); err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &user)
}

func UpdateUser(c *gin.Context) {
	var user entity.UserInfo
	id := c.PostForm("id")
	err := MysqlDB.First(&user, id).Error
	if err != nil {
		c.AbortWithStatus(404)
		log.Println(err.Error())
	} else {
		c.BindJSON(&user)
		MysqlDB.Save(&user)
		c.JSON(http.StatusOK, &user)
	}
}

func ListUser(c *gin.Context) {
	var user []entity.UserInfo
	limit := c.Query("limit")
	MysqlDB.Limit(limit).Find(&user)
	c.JSON(http.StatusOK, &user)
}

func GetUser(c *gin.Context) {
	id := c.Query("id")
	var user entity.UserInfo
	err := MysqlDB.First(&user, id).Error
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user entity.UserInfo
	MysqlDB.Where("id = ?", id).Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"data": "this has been deleted",
	})
}
