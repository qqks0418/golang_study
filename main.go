package main

import (
	"github.com/gin-gonic/gin"
	"github.com/qqks0418/golang_study/controller/env"
)

func main() {

	route := gin.Default()
	v := route.Group("/v1")

	//user.UserApi(v)		// ユーザーAPI
	//task.TaskApi(v)		// タスクAPI
	//todo.TodoApi(); 
	//env.GetEnvApi(v);
	//env.CreateEnvApi(v);
	env.DeleteEnvApi(v);
	//env.EnvHtml(route); 

	// http://localhost:8080/v1/user?category=aaa
	// http://localhost:8080/v1/task/all
	route.Run(":8080")

	// ===============================
	// 確認用
	// ===============================
	//var s int = consts.Add(1,2);
	//d := consts.Hiku(8,1);
	//fmt.Println(s);
	//fmt.Println(d);
}