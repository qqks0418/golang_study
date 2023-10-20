package env

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qqks0418/golang_study/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func DeleteEnvApi(v *gin.RouterGroup) {

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/test_db?parseTime=true")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	boil.SetDB(db)
	//defer db.Close()
	//ctx := context.Background()

	// 削除2
	v.DELETE("/env/:envKey", func(c *gin.Context) {

		envKey := c.Param("envKey")
		fmt.Println(envKey)

		_, err := entity.EnvironmentVariables(
			qm.Where("tenant_id=?", "t7"),
			qm.And("env_key = ?", envKey),
		).DeleteAllG(c)

		if err != nil {
			// ここではエラーを返さない
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{"msg": "成功"})
	})


	// 削除1
	v.DELETE("/env/:envKey/:stageId", func(c *gin.Context) {

		//category := c.DefaultQuery("category", "all")
		
		envKey := c.Param("envKey")
		stageId := c.Param("stageId")

		fmt.Println(envKey)
		fmt.Println(stageId)

		/*
		// 主キーで削除なので条件指定出来ない
		t := &entity.EnvironmentVariable{
			ID: "777",
			TenantID: envKey,
			//EnvKey: "aaa",
			//StageID: "",
		}

		if _, err := t.DeleteG(c); err != nil {
			fmt.Println(err)
			return
		}
		*/

		// DELETE FROM "pilots" WHERE "id"=$1;
		_, err := entity.EnvironmentVariables(
			qm.Where("tenant_id=?", "t7"),
			qm.And("env_key = ?", envKey),
			qm.And("stage_id = ?", stageId),
		).DeleteAllG(c)

		// type safe version of above
		//_, err := entity.EnvironmentVariables(
		//	entity.EnvironmentVariableWhere.TenantID.EQ("t7"),
		//).DeleteAllG(c)

		if err != nil {
			// ここではエラーを返さない
			log.Fatal(err)
		}
		
		c.JSON(http.StatusOK, gin.H{"msg": "成功"})
	})
}