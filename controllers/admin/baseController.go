package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}


func (con BaseController) Success(c *gin.Context, message string, redirectUrl string) {
	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) Error(c *gin.Context, message string, redirectUrl string) {

	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) Response(c *gin.Context, message string,success bool, code int,data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"success": success,
		"code": code,
		"data" : data,
	})
}

func (con BaseController) ResponseNoData(c *gin.Context, message string,success bool, code int) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"success": success,
		"code": code,
		"data" : nil,
	})
}


