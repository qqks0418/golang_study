package env

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qqks0418/golang_study/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// 詳細
var em = map[string][]entity.EnvironmentVariable{}

// レスポンス用
type EnvRes struct {
	EnvList []EnvList `json:"envList"`
	EnvDetailMap map[string][]entity.EnvironmentVariable `json:"envDetailMap"`
}

// リスト用
type EnvList struct {
	Id string `json:"id"`
	EnvKey string `json:"envKey"`
	EnvValue string `json:"envValue"`
}

func GetEnvApi(v *gin.RouterGroup) {

	// データベースのハンドルを取得する
	db, err := sql.Open("mysql", "root:password@(localhost:3306)/test_db?parseTime=true")

	if err != nil {
		// ここではエラーを返さない
		log.Fatal(err)
	}
	boil.SetDB(db)
	//defer db.Close()
	//ctx := context.Background()

	v.GET("/env", func(c *gin.Context) {

		envLists, _ := entity.EnvironmentVariables(
			qm.Where("tenant_id = ?", "t1"),
			//qm.AndIn("env_key in ?", "ek1", "ek2"),
			qm.OrderBy("env_key, default_flg desc"),
		).AllG(c)

		// =====================
		// 一覧・詳細データ加工
		// =====================
		var el []EnvList
		for _, env := range envLists {
			// デフォルトのみ抽出	
			if env.DefaultFLG.Int == 1 {
				el = append(el, EnvList{
					Id: env.ID,
					EnvKey: env.EnvKey.String,
					EnvValue: env.EnvValue.String,
				})
				// 詳細
				detail(c, env.TenantID, env.EnvKey.String)
			}
		}

		// =====================
		// レスポンス設定
		// =====================
		res := EnvRes{}
		res.EnvList = el
		res.EnvDetailMap = em

		c.JSON(http.StatusOK, gin.H{
			"data": res,
		})
	})
}

func detail(c *gin.Context, tenantId string, envKey string) map[string][]entity.EnvironmentVariable {

    envLists, _ := entity.EnvironmentVariables(
		qm.Where("tenant_id = ?", tenantId),
		qm.And("env_key = ?", envKey),
		qm.OrderBy("env_key, default_flg desc"),
	).AllG(c)

	// =====================
	// 詳細用のデータ加工
	// =====================
	var evl = []entity.EnvironmentVariable{}
	for _, env := range envLists {
		evl = append(evl, entity.EnvironmentVariable{
			ID: env.ID,
			TenantID: env.TenantID,
			EnvKey: env.EnvKey,
			StageID: env.StageID,
			EnvValue: env.EnvValue,
			DefaultFLG: env.DefaultFLG,
			UpdatedAt: env.UpdatedAt,
			CreatedAt: env.CreatedAt,
		})
	}
	em[envKey] = evl

    return em
}