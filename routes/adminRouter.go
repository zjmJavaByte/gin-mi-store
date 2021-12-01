package routes

import (
	"gin-mi-store/controllers/admin"
	middlwares2 "gin-mi-store/controllers/admin/middlwares"
	"github.com/gin-gonic/gin"
)

func AdminRouterInit(router *gin.Engine) {
	adminRouter := router.Group("/admin", middlwares2.InitAdminAuthMiddleware)
	{
		adminRouter.GET("/", admin.MainController{}.Index)
		adminRouter.GET("/welcome", admin.MainController{}.Welcome)
		adminRouter.GET("/changeStatus", admin.MainController{}.ChangeStatus)
		adminRouter.GET("/changeNum", admin.MainController{}.ChangeNum)

		adminRouter.GET("/login", admin.LoginController{}.Index)
		adminRouter.GET("/captcha", admin.LoginController{}.Captcha)
		adminRouter.POST("/doLogin", admin.LoginController{}.DoLogin)
		adminRouter.GET("/loginOut", admin.LoginController{}.LoginOut)

		adminRouter.GET("/manager", admin.ManagerController{}.Index)
		adminRouter.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouter.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
		adminRouter.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouter.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
		adminRouter.GET("/manager/delete", admin.ManagerController{}.Delete)

		adminRouter.GET("/focus", admin.FocusController{}.Index)
		adminRouter.GET("/focus/add", admin.FocusController{}.Add)
		adminRouter.POST("/focus/doAdd", admin.FocusController{}.DoAdd)
		adminRouter.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouter.POST("/focus/doEdit", admin.FocusController{}.DoEdit)
		adminRouter.GET("/focus/delete", admin.FocusController{}.Delete)

		adminRouter.GET("/role", admin.RoleController{}.Index)
		adminRouter.GET("/role/add", admin.RoleController{}.Add)
		adminRouter.POST("/role/doAdd", admin.RoleController{}.DoAdd)
		adminRouter.GET("/role/edit", admin.RoleController{}.Edit)
		adminRouter.POST("/role/doEdit", admin.RoleController{}.DoEdit)
		adminRouter.GET("/role/delete", admin.RoleController{}.Delete)
		adminRouter.GET("/role/auth", admin.RoleController{}.Auth)
		adminRouter.POST("/role/doAuth", admin.RoleController{}.DoAuth)

		adminRouter.GET("/access", admin.AccessController{}.Index)
		adminRouter.GET("/access/add", admin.AccessController{}.Add)
		adminRouter.POST("/access/doAdd", admin.AccessController{}.DoAdd)
		adminRouter.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouter.POST("/access/doEdit", admin.AccessController{}.DoEdit)
		adminRouter.GET("/access/delete", admin.AccessController{}.Delete)

		adminRouter.GET("/goodsCate", admin.GoodsCateController{}.Index)
		adminRouter.GET("/goodsCate/add", admin.GoodsCateController{}.Add)
		adminRouter.POST("/goodsCate/doAdd", admin.GoodsCateController{}.DoAdd)
		adminRouter.GET("/goodsCate/edit", admin.GoodsCateController{}.Edit)
		adminRouter.POST("/goodsCate/doEdit", admin.GoodsCateController{}.DoEdit)
		adminRouter.GET("/goodsCate/delete", admin.GoodsCateController{}.Delete)

		adminRouter.GET("/goodsType", admin.GoodsTypeController{}.Index)
		adminRouter.GET("/goodsType/add", admin.GoodsTypeController{}.Add)
		adminRouter.POST("/goodsType/doAdd", admin.GoodsTypeController{}.DoAdd)
		adminRouter.GET("/goodsType/edit", admin.GoodsTypeController{}.Edit)
		adminRouter.POST("/goodsType/doEdit", admin.GoodsTypeController{}.DoEdit)
		adminRouter.GET("/goodsType/delete", admin.GoodsTypeController{}.Delete)

		adminRouter.GET("/goodsTypeAttribute", admin.GoodsTypeAttributeController{}.Index)
		adminRouter.GET("/goodsTypeAttribute/add", admin.GoodsTypeAttributeController{}.Add)
		adminRouter.POST("/goodsTypeAttribute/doAdd", admin.GoodsTypeAttributeController{}.DoAdd)
		adminRouter.GET("/goodsTypeAttribute/edit", admin.GoodsTypeAttributeController{}.Edit)
		adminRouter.POST("/goodsTypeAttribute/doEdit", admin.GoodsTypeAttributeController{}.DoEdit)
		adminRouter.GET("/goodsTypeAttribute/delete", admin.GoodsTypeAttributeController{}.Delete)

		adminRouter.GET("/goods", admin.GoodsController{}.Index)
		adminRouter.GET("/goods/add", admin.GoodsController{}.Add)
		adminRouter.POST("/goods/doAdd", admin.GoodsController{}.DoAdd)
		adminRouter.GET("/goods/edit", admin.GoodsController{}.Edit)
		adminRouter.POST("/goods/doEdit", admin.GoodsController{}.DoEdit)
		adminRouter.GET("/goods/changeGoodsImageColor", admin.GoodsController{}.ChangeGoodsImageColor)
		adminRouter.GET("/goods/removeGoodsImage", admin.GoodsController{}.RemoveGoodsImage)
		adminRouter.GET("/goods/goodsTypeAttribute", admin.GoodsController{}.GoodsTypeAttribute)
		adminRouter.POST("/goods/imageUpload", admin.GoodsController{}.ImageUpload)
		adminRouter.GET("/goods/delete", admin.GoodsController{}.Delete)

		adminRouter.GET("/nav", admin.NavController{}.Index)
		adminRouter.GET("/nav/add", admin.NavController{}.Add)
		adminRouter.POST("/nav/doAdd", admin.NavController{}.DoAdd)
		adminRouter.GET("/nav/edit", admin.NavController{}.Edit)
		adminRouter.POST("/nav/doEdit", admin.NavController{}.DoEdit)
		adminRouter.GET("/nav/delete", admin.NavController{}.Delete)

		adminRouter.GET("/setting", admin.SettingController{}.Index)
		adminRouter.POST("/setting/doEdit", admin.SettingController{}.DoEdit)
	}
}
