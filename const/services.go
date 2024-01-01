package _const

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strconv"
	"time"
)

var Services *Service

/**
* Services instructor
 */
type Service struct {
	GIN                          *gin.Context
	DB                           *gorm.DB
	DBERR                        error
	Modules                      map[string]map[string]interface{}
	TimeLocation                 *time.Location
	Time                         time.Time
	RequestBody                  []byte
	GinCopy                      *gin.Context
	SQL                          *sql.DB
	SQLX                         *sqlx.DB
	SupportedLanguageCount       int
	SupportedLanguageCountString string
	SupportedLanguageMap         map[string]string
}

/**
* Init Services
 */
func ServicesInit() *Service {
	length := len(SupportedLang())
	return &Service{SupportedLanguageCount: length,
		SupportedLanguageCountString: strconv.Itoa(length),
		SupportedLanguageMap:         SupportedLang(),
	}
}

/**
* short hand to get request
 */
func Request() *http.Request {
	return Services.GinCopy.Request
}

/**
* short hand to get body
* with []byte
 */
func GetBodyB() []byte {
	return Services.RequestBody
}

/**
* short hand to get body
* as string
 */
func GetBody() string {
	return string(GetBodyB())
}
