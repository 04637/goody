package router

import (
	"aid.dev/goody/db"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	// 路径映射
	router.POST("/user/create", db.CreateUser)
	router.GET("/user/list", db.ListUser)
	router.PUT("/user/update", db.UpdateUser)
	router.GET("/user/find", db.GetUser)
	router.DELETE("/user/:id", db.DeleteUser)
	_ = router.Run(":8080")
}
