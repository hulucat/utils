package utils

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/cihub/seelog"
	"github.com/gorilla/sessions"
	"github.com/hulucat/conf"
	"os"
)

var logger seelog.LoggerInterface

func Debug(content string) {
	logger.Debug(content)
}

func Info(content string) {
	logger.Debug(content)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func init() {
	appConf := `
		<seelog type="sync">
			<outputs>
				<rollingfile type="date" filename="./logs/` + os.Args[0] + `.log" datepattern="20060102" formatid="main"/>
			</outputs>
			<formats>
				<format id="main" format="%Date %Time [%LEV] | %Msg%n"/>
			</formats>
		</seelog>`
	var err error
	logger, err = seelog.LoggerFromConfigAsBytes([]byte(appConf))
	if err != nil {
		fmt.Errorf("Error loading log configuration: %s \n", err.Error())
		panic(err)
	}
	Mclient = memcache.New(conf.Get("memcached"))
	sessionStore = sessions.NewCookieStore([]byte(conf.Get("cookie_secret")))
}
