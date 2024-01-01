package auth

import (
	"encoding/json"
	"github.com/bykovme/gotrans"
	"github.com/gin-gonic/gin"
	"os"
	"starter-golang-new/app/models"
	"starter-golang-new/app/requests/visitor"
	"starter-golang-new/app/transformers"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
)

/**
* check if user have access to login in system
 */
func Login(g *gin.Context) {
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Login()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	request := models.ConvertBodyToHashMap()
	user, valid := checkUserExistsNotBlocked(request["email"], "")
	if !valid {
		return
	}
	/**
	* now check if password are valid
	* if user password is not valid we will return invalid email
	* or password
	 */
	check := helpers.CheckPasswordHash(request["password"], user.Password)
	if !check {
		helpers.ReturnNotFoundUser(gotrans.Tr(helpers.GetCurrentLang(), "not_valid_password"))
		return
	}
	// update token then return with the new data
	token, _ := helpers.GenerateToken(user.Password + user.Email)
	_const.Services.DB.Model(&user).Update("token", token).First(&user)

	helpers.OkResponse(gotrans.Tr(helpers.GetCurrentLang(), "login"), returnUser(user))
}

/**
* Register new user on system
 */
func Register(g *gin.Context) {
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Register()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return
	}

	/**
	* set role and block
	* role 1 is user
	* block user (1 , 2) 2 is not block 1 is block
	 */
	request := models.ConvertBodyToHashMap()
	user := models.User{
		Name:              request["name"].(string),
		Email:             request["email"].(string),
		PermissionGroupId: _const.USER_ID,
		Password:          request["password"].(string),
		Status:            _const.ACTIVE,
	}
	/**
	* create new user based on register struct
	* token , role  , block will set with event
	 */
	_const.Services.DB.Create(&user)
	// now user is login we can return his info
	helpers.OkResponse(gotrans.Tr(helpers.GetCurrentLang(), "success_register"), returnUser(user))
}

/**
* recover password take request token
* select user that have this token
* if user token valid and user not block
* then user can  recover his password
 */
func Recover(g *gin.Context) {
	//init Reset struct to validate request
	recoverPassword := new(models.Recover)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Recover()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return
	}

	var d = make(map[string]interface{})
	body := _const.GetBodyB()
	_ = json.Unmarshal(body, &d)
	if _, ok := d["token"]; ok {
		if d["token"].(string) != "" {
			recoverPassword.Token = d["token"].(string)
		}
	}
	if _, ok := d["password"]; ok {
		if d["password"].(string) != "" {
			recoverPassword.Password = d["password"].(string)
		}
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	user, valid := checkUserExistsNotBlocked("", recoverPassword.Token)
	if !valid {
		return
	}
	/**
	* now update token and update password
	* we update token to make it the old link not valid
	 */
	user.Password = recoverPassword.Password
	encPassword, _ := helpers.HashPassword(user.Password)
	token, _ := helpers.GenerateToken(user.Password + user.Email)
	_const.Services.DB.Model(&user).Updates(map[string]interface{}{"password": encPassword, "token": token})
	user.Token = token
	// notice user that his password has been changes
	sendRecoverPasswordEmail(user)
	// return ok response
	helpers.OkResponse(gotrans.Tr(helpers.GetCurrentLang(), "reset_password"), transformers.UserResponse(user))
}

/***
* notice user that his password has been updated
 */
func sendRecoverPasswordEmail(user models.User) {
	msg := "<h6>Your Password has been updated to (" + user.Password + ")</h6>" + "<br>"
	msg += "<h6>Do not worry your password is encrypted , this just note for your activity</h6>"
	helpers.SendMail(user.Email, "Your password has been updated", msg)
}

/**
* reset password
* with email you can send reset link
* to user email
 */
func Reset(g *gin.Context) {
	// init Reset struct to validate request
	reset := new(models.Reset)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := visitor.Reset()
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err) {
		return
	}
	var d = make(map[string]interface{})
	body := _const.GetBodyB()
	_ = json.Unmarshal(body, &d)
	if _, ok := d["email"]; ok {
		if d["email"].(string) != "" {
			reset.Email = d["email"].(string)
		}
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	user, valid := checkUserExistsNotBlocked(reset.Email, "")
	if !valid {
		return
	}
	sendRestLink(user)
	// return ok response
	var data map[string]interface{}
	helpers.OkResponse(gotrans.Tr(helpers.GetCurrentLang(), "reset_password_link"), data)
}

/**
* create reset password link
* send it to user email
 */
func sendRestLink(user models.User) {
	msg := "<h6> Your Request To reset your password if you take this action click on this link to reset your password </h6>" + "<br>"
	msg += `<a href="` + os.Getenv("RESET_PASSWORD_URL") + user.Token + `">Reset Password</a>`
	helpers.SendMail(user.Email, "Reset Password Request", msg)
}

/**
* check if user exists
* check if user not blocked
 */
func checkUserExistsNotBlocked(email interface{}, token interface{}) (models.User, bool) {
	// init user struct binding data for user
	var user models.User
	/**
	* check if this email exists database
	* if this email will not found will return not found
	* will return 404 code
	* will select by email if token is empty
	* if token not empty select by token
	 */
	if token != "" {
		_const.Services.DB.Find(&user, "token = ? ", token)
	} else {
		_const.Services.DB.Find(&user, "email = ? ", email)
	}
	if user.ID == 0 {
		helpers.ReturnNotFoundUser(gotrans.Tr(helpers.GetCurrentLang(), "not_found_user"))
		return user, false
	}
	// if user block
	if user.Status != _const.ACTIVE {
		helpers.ReturnForbidden(gotrans.Tr(helpers.GetCurrentLang(), "blocked"))
		return user, false
	}
	var permissionGroup models.PermissionGroup
	// check if this id exits , abort if not
	if models.FindOrFail(user.PermissionGroupId, &permissionGroup); permissionGroup.ID == 0 {
		helpers.ReturnForbidden(gotrans.Tr(helpers.GetCurrentLang(), "not_found_permission"))
		return user, false
	}
	/// check if permission group active
	if permissionGroup.Status != _const.ACTIVE {
		helpers.ReturnForbidden(gotrans.Tr(helpers.GetCurrentLang(), "permission_not_Active"))
		return user, false
	}
	return user, true
}

/***
* get user by token
* get user data from header
* return user data
 */
func GetUserByToken(g *gin.Context) {
	user := models.Auth()
	if user.ID == 0 {
		helpers.ReturnYouAreNotAuthorize()
		g.Abort()
		return
	}
	// now return row data after transformers
	helpers.OkResponse(helpers.DoneGetItem(), returnUser(user))
}

/**
* stander response user with permissions
 */
func returnUser(user models.User) map[string]interface{} {
	// get user role
	var roles []models.Role
	_const.Services.DB.Where("permission_group_id = ?", user.PermissionGroupId).Find(&roles)
	/// build response
	data := make(map[string]interface{})
	data["user"] = transformers.UserResponse(user)
	data["role"] = transformers.RolesResponse(roles)

	return data
}
