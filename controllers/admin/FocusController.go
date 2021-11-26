package admin

import (
	"gin-mi-store/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	var focus []models.Focus
	models.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focus,
	})

}
func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {

	title := c.PostForm("title")
	focusType,_ := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort,_ := strconv.Atoi(c.PostForm("sort"))
	status,_ := strconv.Atoi(c.PostForm("status"))
	//file, err := c.FormFile("focus_img")
	focusImg, err := models.UploadImg(c, "focus_img")
	if err != nil {
		con.Error(c, err.Error(), "/admin/focus/add")
		return
	}
	/*ext := path.Ext(file.Filename)
	fileType := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if b := fileType[ext]; !b {
		con.Error(c, "文件类型不支持", "/admin/focus/add")
		return
	}
	day := models.GetDay()
	dir := "static/upload/" + day
	if err = os.MkdirAll(dir, 0666);err != nil{
		con.Error(c, err.Error(), "/admin/focus/add")
		return
	}
	formatInt := strconv.FormatInt(models.GetUnix(), 10)
	savePath := path.Join(dir, formatInt+ext)
	if err = c.SaveUploadedFile(file, savePath);err != nil{
		con.Error(c, err.Error(), "/admin/focus/add")
		return
	}*/
	focus := models.Focus{
		AddTime:   int(models.GetUnix()),
		Title:     title,
		FocusType: focusType,
		Link:      link,
		Sort:      sort,
		Status:    status,
		FocusImg:  focusImg,
	}
	err = models.DB.Create(&focus).Error
	if err != nil {
		con.Error(c, "数据保存错误", "/admin/focus/add")
		return
	}
	con.Success(c,"保存成功","/admin/focus")
}

func (con FocusController) Edit(c *gin.Context) {
	id,err := models.Int(c.Query("id"))
	if err != nil {
		con.Error(c, "传入数据错误", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus" : focus,
	})
}
func (con FocusController) Delete(c *gin.Context) {
	id, err1 := models.Int(c.Query("id"))
	if err1 != nil {
		con.Error(c, "非法请求", "/admin/focus")
		return
	}
	focus := models.Focus{Id: id}
	models.DB.Delete(&focus)
	//根据自己的需要 要不要删除图片
	// os.Remove("static/upload/20210915/1631694117.jpg")
	con.Success(c,"成功","/admin/focus")
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := models.Int(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := models.Int(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := models.Int(c.PostForm("sort"))
	status, err4 := models.Int(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.Error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.Error(c, "请输入正确的排序值", "/admin/focus/edit?id="+models.String(id))
	}
	//上传文件
	focusImg, _ := models.UploadImg(c, "focus_img")

	focus := models.Focus{Id: id}
	models.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}
	err5 := models.DB.Save(&focus).Error
	if err5 != nil {
		con.Error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+models.String(id))
	} else {
		con.Success(c, "增加轮播图成功", "/admin/focus")
	}
}
