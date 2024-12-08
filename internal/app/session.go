package app

import (
	"time"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

var sessMgr *scs.SessionManager

func InitSessionManager() {
	sessMgr = scs.New()

	if GetEnv().AppEnv == Development {
		sqlDB := GetSQLDB()
		sessMgr.Store = postgresstore.New(sqlDB)
	}

	sessMgr.Lifetime = 24 * 7 * time.Hour
	if GetEnv().AppEnv == Production {
		sessMgr.Cookie.Secure = true
	}
}

func GetSession() *scs.SessionManager {
	return sessMgr
}
