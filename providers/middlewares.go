package providers

import (
	"fmt"
	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"os"
	"starter-golang-new/app/middleware"
	"strconv"
)

func middlewares(r *gin.Engine) *gin.Engine {
	/// run cors middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Language())
	r.Use(middleware.Time())
	limitNumber, _ := strconv.Atoi(os.Getenv("MAX_ALLOWED_REQUESTS"))
	r.Use(limit.MaxAllowed(limitNumber))
	/// allow dump from env file
	allowDump, _ := strconv.ParseBool(os.Getenv("ALLOW_DUMP_GIN"))
	if allowDump {
		//use Dump() default will print on stdout
		r.Use(gindump.Dump())
		//or use DumpWithOptions() with more options
		r.Use(gindump.DumpWithOptions(true, true, false, true, false, func(dumpStr string) {
			fmt.Println(dumpStr)
		}))
	}

	return r
}
