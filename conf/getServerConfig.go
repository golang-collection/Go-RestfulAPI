package watchConfig

import (
	"github.com/spf13/viper"
	"superTools-backend/services/watchConfig/model"
)

/**
* @Author: super
* @Date: 2020-08-21 21:19
* @Description:
**/

var Server = &model.ServerConfig{}


func GetServerConfig() {
	httpPort := viper.GetInt("server.http_port")
	readTimeout := viper.GetDuration("server.read_timeout")
	writeTimeout := viper.GetDuration("server.write_timeout")

	Server.HttpPort = httpPort
	Server.ReadTimeout = readTimeout
	Server.WriteTimeout = writeTimeout
}
