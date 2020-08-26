package logging

import (
	"fmt"
	"github.com/spf13/viper"
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
	return fmt.Sprintf("%s%s", viper.GetString("log.runtime_root_path"), viper.GetString("log.log_save_path"))
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		viper.GetString("log.log_save_name"),
		time.Now().Format(viper.GetString("log.time_format")),
		viper.GetString("log.log_file_ext"),
	)
}
