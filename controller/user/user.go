package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name"`
}

func UserApi(v *gin.RouterGroup) {

	v.GET("/user", func(c *gin.Context) {
		//category := c.Query("category")
		// パラメータがnilの時は第二引数の値になる
		category := c.DefaultQuery("category", "all")
		fmt.Println(category);

		//リダイレクト
		//c.Redirect(http.StatusMovedPermanently, "https://www.google.com/")

		c.JSON(200, gin.H{
			"id": 1,
			"name":"ユーザー",
			"map": map[string]int{
				"key1": 100,
				"key2": 200,
			},
		})
	})

	v.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// データベースに登録

		c.JSON(201, gin.H{"msg": "", "data":user})

	})

	v.PUT("/user/:userID", func(c *gin.Context) {
		userID := c.Param("userID")
		
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// データベースを更新

		c.JSON(200, gin.H{"msg": "OK_OK_OK", "id":userID})
	})

	v.DELETE("/user/:userID", func(c *gin.Context) {
		deleteUserID := c.Param("userID")

		// データベースから削除

		c.JSON(200, gin.H{"msg": "OK_OK_OK", "id":deleteUserID})
	})
}