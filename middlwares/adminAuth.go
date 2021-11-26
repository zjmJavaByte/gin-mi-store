package middlwares

import (
	"encoding/json"
	"fmt"
	"gin-mi-store/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"os"
	"strings"
)


func InitAdminAuthMiddleware(c *gin.Context) {
	//进行权限判断 没有登录的用户 不能进入后台管理中心

	//1、获取Url访问的地址  /admin/captcha

	//2、获取Session里面保存的用户信息

	//3、判断Session中的用户信息是否存在，如果不存在跳转到登录页面（注意需要判断） 如果存在继续向下执行

	//4、如果Session不存在，判断当前访问的URl是否是login doLogin captcha，如果不是跳转到登录页面，如果是不行任何操作

	pathname := strings.Split(c.Request.URL.String(), "?")[0]

	session := sessions.Default(c)
	userinfo := session.Get("userinfo")
	userinfoStr,ok := userinfo.(string)
	if ok {
		var userinfoStruct []models.Manager
		err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
		if err != nil || !(len(userinfoStruct) > 0 && userinfoStruct[0].Username != "") {
			if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
				c.Redirect(302, "/admin/login")
			}
		}else {
			urlPath := strings.Replace(pathname, "/admin/", "", 1)
			if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/" + urlPath) {
				roleId := userinfoStruct[0].RoleId
				var roleAccess []models.RoleAccess
				models.DB.Where("role_id = ?",roleId).Find(&roleAccess)
				m := make(map[int]bool)
				for _, access := range roleAccess {
					m[access.AccessId] = true
				}
				access := models.Access{}
				models.DB.Where("url = ?",urlPath).Find(&access)
				if b := m[access.Id];!b {
					c.String(200, "没有权限")
					c.Abort()
				}
			}
		}
	}else {
		//没有登录
		if pathname != "/admin/login" && pathname != "/admin/doLogin" && pathname != "/admin/captcha" {
			c.Redirect(302, "/admin/login")
		}
	}

}


//排除权限判断的方法

func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("conf/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}
	excludeAuthPath := config.Section("").Key("excludeAuthPath").String()

	excludeAuthPathSlice := strings.Split(excludeAuthPath, ",")
	// return true
	fmt.Println(excludeAuthPathSlice)
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}