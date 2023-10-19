package env

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func EnvHtml(r *gin.Engine) {
	r.LoadHTMLFiles("./tmpl/index.html")
	r.GET("/env", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{})
	})
}