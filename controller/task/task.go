package task

import (
	"github.com/gin-gonic/gin"
)

type Task struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name"`
	TaskMap map[int]string `json:"taskMap"`
	ToolItem []string `json:"toolItem"`
}

func TaskApi(v *gin.RouterGroup) {

	v.GET("/task/all", func(c *gin.Context) {
		category := c.DefaultQuery("category", "all")

		// 構造体に追加ORMがやる予定のこと
		m := map[int]string{1: "Next.js", 2: "Go言語"}
		if category == "all" {
			m[3] = "Java"
		}
		//toolItem := [...]string{"vsCode", "git"}
		var toolItem []string
		toolItem = append(toolItem, "vsCode")
		toolItem = append(toolItem, "git")
		res := Task{Id: 1, Name: "タスク一覧", TaskMap: m, ToolItem: toolItem}

		c.JSON(200, gin.H{
			"data": res,
		})
	})
}