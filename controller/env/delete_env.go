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
	boil.DebugMode = true
	//defer db.Close()
	//ctx := context.Background()

	v.DELETE("/env/:name", func(c *gin.Context) {

		name := c.Param("name")
		fmt.Println(name)

		ess, _ := entity.EnvSettings(
			qm.Where("tenant_id = ?", "A"),
			qm.And("name = ?", name),
		).AllG(c)

		tagetDels := [] string{"11", "12"}

		for _, v := range ess {
			fmt.Println(v.ID)
			tagetDels = append(tagetDels, v.ID)
		}

		tagets := make([]interface{}, len(tagetDels))
		for i, taget := range tagetDels {
			tagets[i] = taget
		}

		_, err := entity.Envs(
			qm.WhereIn("env_setting_id IN ?", tagets...),
		).DeleteAllG(c)

		if err != nil {
			// ここではエラーを返さない
			log.Fatal(err)
		}

		_, es_err := entity.EnvSettings(
			qm.Where("tenant_id = ?", "A"),
			qm.And("name = ?", name),
		).DeleteAllG(c)

		if es_err != nil {
			// ここではエラーを返さない
			log.Fatal(es_err)
		}

		stages, _ := entity.Stages(
			qm.Where("tenant_id = ?", "A"),
		).AllG(c)

		for _, v := range stages {
			fmt.Println(v.StageName)
		}

		c.JSON(http.StatusOK, gin.H{"msg": "成功"})
	})

	/*
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
	*/


	/*
	// 削除1
	v.DELETE("/env/:envKey/:stageId", func(c *gin.Context) {

		//category := c.DefaultQuery("category", "all")
		
		envKey := c.Param("envKey")
		stageId := c.Param("stageId")

		fmt.Println(envKey)
		fmt.Println(stageId)

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
	*/
}