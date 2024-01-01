package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
	_const "starter-golang-new/const"
	"starter-golang-new/helpers"
	"time"
)

/**
* set time zone
* if fails we will call the default time zone form the env file
 */
func Time() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeZone := c.GetHeader("TimeZone")
		if timeZone == "" {
			timeZone = os.Getenv("TIME_ZONE")
		}
		timeZone = helpers.ClearText(timeZone)
		loc, err := time.LoadLocation(timeZone)
		var newTime time.Time
		if err == nil {
			_const.Services.Time = newTime.In(loc)
		} else {
			loc, _ = time.LoadLocation(os.Getenv("TIME_ZONE"))
			_const.Services.Time = newTime.In(loc)
		}
		_const.Services.TimeLocation = loc
	}
}
