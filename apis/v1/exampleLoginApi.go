package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/params"
	"github.com/jangozw/gin-api-common/services"
	"github.com/jangozw/gin-api-common/utils"
)

//login api

func Login(c *gin.Context) {
	p := params.Login{}
	if err := c.ShouldBind(&p); err != nil {
		utils.Ctx(c).Fail(err)
		return
	}
	jwtToken, err := services.AppLogin(p.Mobile, p.Pwd)
	if err != nil {
		utils.Ctx(c).Fail(err)
		return
	}
	utils.Ctx(c).Success(map[string]interface{}{"token": jwtToken})
	return
}

//logout api
func Logout(c *gin.Context) {
	userId := utils.Ctx(c).GetLoginUid()
	if err := services.AppLogout(userId); err != nil {
		utils.Ctx(c).Fail(err)
		return
	}
	utils.Ctx(c).SuccessSimple()
	return
}
