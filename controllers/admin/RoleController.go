package admin

import (
	"fmt"
	"gin-mi-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(ctx *gin.Context) {
	var roles []models.Role
	models.DB.Find(&roles)
	ctx.HTML(http.StatusOK,"admin/role/index.html",gin.H{
		"roleList" : roles,
	})
}

func (con RoleController) Add(ctx *gin.Context) {
	ctx.HTML(http.StatusOK,"admin/role/add.html",gin.H{

	})
}

func (con RoleController) DoAdd(ctx *gin.Context) {
	title := ctx.PostForm("title")
	description := ctx.PostForm("description")
	status,_ := models.Int(ctx.PostForm("Status"))
	role := models.Role{
		Title:       title,
		Description: description,
		AddTime:     int(models.GetUnix()),
		Status:      status,
	}
	if err := models.DB.Create(&role).Error;err != nil{
		con.Error(ctx,err.Error(),"/admin/role/add")
		return
	}
	con.Success(ctx,"保存成功","/admin/role")

}

func (con RoleController) Edit(ctx *gin.Context) {
	id, err := models.Int(ctx.Query("id"))
	if err != nil {
		con.Error(ctx,err.Error(),"/admin/role")
		return
	}
	role := models.Role{
		Id: id,
	}
	if err := models.DB.Find(&role).Error;err != nil{
		con.Error(ctx,err.Error(),"/admin/role")
		return
	}
	ctx.HTML(http.StatusOK,"admin/role/edit.html",gin.H{
		"role" : role,
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	status, err2 := models.Int(c.PostForm("Status"))
	if err1 != nil || err2 != nil {
		con.Error(c, "参数错误", "/admin/role")
		return
	}
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")

	if title == "" {
		con.Error(c, "角色的标题不能为空", "/admin/role/edit?id="+models.String(id))
		return
	}

	role := models.Role{Id: id}
	models.DB.Find(&role)
	role.Title = title
	role.Description = description
	role.Status = status
	err2 = models.DB.Save(&role).Error
	if err2 != nil {
		con.Error(c, "修改数据失败", "/admin/role/edit?id="+models.String(id))
	} else {
		con.Success(c, "修改数据成功", "/admin/role")
	}

	//查询要修改的数据 然后 修改

	// c.String(http.StatusOK, "-执行修改")
}

func (con RoleController) Delete(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
	} else {
		role := models.Role{Id: id}
		models.DB.Delete(&role)
		con.Success(c, "删除数据成功", "/admin/role")
	}
}

func (con RoleController) Auth(c *gin.Context) {
	id, err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role")
		return
	}
	var accessList []models.Access
	models.DB.Where("module_id=?",0).Preload("AccessItem").Find(&accessList)

	var roleAccess []models.RoleAccess
	models.DB.Where("role_id = ?",id).Find(&roleAccess)

	m := make(map[int]bool)
	for _, access := range roleAccess {
		m[access.AccessId] = true
	}

	for i, access := range accessList {
		if b := m[access.Id]; b {
			accessList[i].Checked = true
		}
		for j, a := range access.AccessItem {
			if b1 := m[a.Id]; b1 {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}
	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     id,
		"accessList": accessList,
	})

}

func (con RoleController) DoAuth(c *gin.Context) {
	roleId, err := models.Int(c.PostForm("role_id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/role/auth?id="+models.String(roleId))
		return
	}
	var access models.RoleAccess
	models.DB.Where("role_id=?", roleId).Delete(&access)
	accessIds := c.PostFormArray("access_node[]")
	for _, id := range accessIds {
		accessId,_ := models.Int(id)
		roleAccess := models.RoleAccess{
			RoleId:   roleId,
			AccessId: accessId,
		}
		models.DB.Create(&roleAccess)
	}
	fmt.Println(roleId)
	fmt.Println(accessIds)

	fmt.Println("/admin/role/auth?id=?" + models.String(roleId))
	// c.String(200, "DoAuth")
	// admin/role/auth?id=9
	con.Success(c, "授权成功", "/admin/role/auth?id="+models.String(roleId))

}
