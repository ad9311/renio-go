package conf

import (
	"time"

	"github.com/alexedwards/scs/v2"
)

var sessMgr *scs.SessionManager

func InitSessionManager() {
	sessMgr = scs.New()

	sessMgr.Lifetime = 24 * 7 * time.Hour
	if GetEnv().AppEnv == Production {
		sessMgr.Cookie.Secure = true
	}
}

func GetSession() *scs.SessionManager {
	return sessMgr
}
