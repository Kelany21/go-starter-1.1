package middleware

import (
	"encoding/json"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"starter-golang-new/app/models"
	"starter-golang-new/config"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"strconv"
)

/**
* This middle ware will Allow only
* Will block not admin role admin role is (2)
* if user allow to access then this middleware will add
* one header with user information you can use later (ADMIN_DATA)
* in function you call
 */
func Admin() gin.HandlerFunc {
	return func(g *gin.Context) {
		functionName := helpers.GetMethodName(g.HandlerName())
		//var user models.User
		/// get Authorization header to check if user send it first
		adminToken := helpers.GetClearToken()
		if adminToken == "" {
			helpers.ReturnYouAreNotAuthorize()
			g.Abort()
			return
		}
		/// check if token exits in database
		user := models.Auth()
		if user.ID == 0 {
			helpers.ReturnYouAreNotAuthorize()
			g.Abort()
			return
		}
		/// check if user block or not
		if user.Status != _const.ACTIVE {
			helpers.ReturnYouAreNotAuthorize()
			g.Abort()
			return
		}
		if _, ok := config.GlobalPermissions()[functionName]; !ok {
			var role models.Role
			_const.Services.DB.Where("func_name = ? AND permission_group_id = ? ", functionName, user.PermissionGroupId).First(&role)
			if role.ID == 0 {
				helpers.ReturnForbidden(gotrans.Tr(helpers.GetCurrentLang(), "403"))
				g.Abort()
				return
			}
		}
		/// not set header with user information
		userJson, _ := json.Marshal(&user)
		g.Request.Header.Set("ADMIN_DATA", string(userJson))
		g.Request.Header.Set("user_id", strconv.Itoa(int(user.ID)))
		g.Next()
	}
}
