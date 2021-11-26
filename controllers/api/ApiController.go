package api

import (
	"gin-mi-store/controllers/admin"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	admin.BaseController
}

func (con ApiController) Index(c *gin.Context) {

}

func (con ApiController) Userlist(c *gin.Context) {

}

func (con ApiController) Plist(c *gin.Context) {

}
