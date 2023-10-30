package env

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qqks0418/golang_study/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateEnvApi(v *gin.RouterGroup) {

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/test_db?parseTime=true")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	boil.SetDB(db)
	//defer db.Close()
	//ctx := context.Background()

	v.POST("/env", func(c *gin.Context) {
		// リクエスト取得
		var reqEvList []entity.EnvironmentVariable
		if err := c.ShouldBindJSON(&reqEvList); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("---- " + reqEvList[0].EnvKey.String)

		// トランザクション（例）
		/*
		tx, err := boil.BeginTx(c, nil)
		if err != nil {
			return
		}
		var tt entity.EnvironmentVariable
		tt.ID = "aaaa"
		if err := tt.Upsert(c, tx, boil.Whitelist("env_value"), boil.Infer()); err != nil {
			tx.Rollback()
			return
		}

		if err := tx.Commit(); err != nil {
			tx.Rollback()
			panic(err)
		}
		*/

		// データベースに登録
		for _, ev := range reqEvList {
			//if ev.EnvValue.String == "" {
			//	continue
			//}
			var t entity.EnvironmentVariable
			t.ID = ev.ID
			t.TenantID = ev.TenantID
			t.EnvKey = ev.EnvKey
			t.StageID = ev.StageID
			t.EnvValue = ev.EnvValue
			t.DefaultFLG = ev.DefaultFLG 

			// 登録・更新
			// 3番目引数 UPDATEフィールド, 4番目の引数 INSERTフィールド
			err := t.Upsert(c, db, boil.Whitelist("env_value"), boil.Infer())

			if err != nil {
				// ここではエラーを返さない
				log.Fatal(err)
			}
			

			//if err := t.Upsert(c, db, boil.Whitelist("env_value"), boil.Infer()); err != nil {
			//	fmt.Println(err)
			//}
		}
		c.JSON(201, gin.H{"msg": "成功", "data":reqEvList})
	})
}