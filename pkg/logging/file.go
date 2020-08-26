package logging

import (
	"Go-RestfulAPI/conf"
	"fmt"
	"time"
)

/**
* @Author: super
* @Date: 2020-08-26 15:00
* @Description:
**/

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s%s", watchConfig.LogConfig.RuntimeRootPath, watchConfig.LogConfig.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		watchConfig.LogConfig.LogSaveName,
		time.Now().Format(watchConfig.LogConfig.TimeFormat),
		watchConfig.LogConfig.LogFileExt,
	)
}
