package watchConfig

import (
	"Go-RestfulAPI/conf/model"
	"fmt"
	"github.com/spf13/viper"
)

/**
* @Author: super
* @Date: 2020-08-21 19:58
* @Description:
**/

var Mysql = &model.MySQLConfig{}
var Redis = &model.RedisConfig{}
var Rabbitmq = &model.RabbitMQConfig{}
var LogConfig = &model.LogConfig{}
var RunMode string
var JwtSecret string

func init() {
	viper.SetConfigFile("services/watchConfig/config/config.json") //文件名
	err := viper.ReadInConfig()                                    // 会查找和读取配置文件
	if err != nil {
		// Handle errors reading the config file
		fmt.Printf("read config error: %v\n", err)
		panic(err)
	}

	GetRedisUrl()
	GetRabbitMQUrl()
	GetMysqlUrl()
	GetJwtSecret()
	GetLogConfig()
	GetRunMode()
	GetServerConfig()
}

func GetMysqlUrl() {
	mysqlHost := viper.GetString("mysql.host")
	mysqlUser := viper.GetString("mysql.user")
	mysqlPassword := viper.GetString("mysql.password")
	mysqlDBName := viper.GetString("mysql.db_name")
	mysqlURL := mysqlUser + ":" + mysqlPassword + "@(" + mysqlHost + ")/" + mysqlDBName

	Mysql.User = mysqlUser
	Mysql.Password = mysqlPassword
	Mysql.Host = mysqlHost
	Mysql.DBName = mysqlDBName
	Mysql.URL = mysqlURL
}

func GetRedisUrl() {
	redisURL := viper.GetString("redis.host")
	maxIdle := viper.GetInt("redis.maxIdle")
	max_activate := viper.GetInt("redis.max_activate")
	Redis.Url = redisURL
	Redis.MaxActive = max_activate
	Redis.MaxIdle = maxIdle
}

func GetRabbitMQUrl() {
	mqHost := viper.GetString("rabbitmq.host")
	mqUser := viper.GetString("rabbitmq.user")
	mqPassword := viper.GetString("rabbitmq.password")
	mqURL := "amqp://" + mqUser + ":" + mqPassword + "@" + mqHost + "/"

	Rabbitmq.Host = mqHost
	Rabbitmq.User = mqUser
	Rabbitmq.Password = mqPassword
	Rabbitmq.URL = mqURL

}

func GetRunMode() {
	RunMode = viper.GetString("run_mode")
}

func GetJwtSecret() {
	JwtSecret = viper.GetString("jwt_secret")
}

func GetLogConfig() {
	logSavePath := viper.GetString("log.log_save_path")
	logSaveName := viper.GetString("log.log_save_name")
	logFileExt := viper.GetString("log.log_file_ext")
	timeFormat := viper.GetString("log.time_format")
	runtimePath := viper.GetString("log.runtime_root_path")

	LogConfig.LogSavePath = logSavePath
	LogConfig.LogSaveName = logSaveName
	LogConfig.LogFileExt = logFileExt
	LogConfig.TimeFormat = timeFormat
	LogConfig.RuntimeRootPath = runtimePath
}
