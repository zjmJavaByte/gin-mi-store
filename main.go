package main

import (
	"gin-mi-store/models"
	"gin-mi-store/routes"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {

	g := gin.Default()

	//自定义模板函数  注意要把这个函数放在加载模板前
	g.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"Str2Html":   models.Str2Html,
		"FormatImg":  models.FormatImg,
		"Sub":        models.Sub,
		"Substr":     models.Substr,
	})

	//加载模板 放在配置路由前面
	g.LoadHTMLGlob("templates/**/**/*")
	//配置静态web目录   第一个参数表示路由, 第二个参数表示映射的目录
	g.Static("/static", "./static")
	// 创建基于 cookie 的存储引擎，secret11111 参数是用于加密的密钥

	store := cookie.NewStore([]byte("secret111"))
	//配置session的中间件 store是前面创建的存储引擎，我们可以替换成其他存储引擎
	g.Use(sessions.Sessions("mysession", store))

	routes.AdminRouterInit(g)
	routes.ApiRoutersInit(g)
	routes.DefaultRoutersInit(g)
	g.Run(":8088")

}
