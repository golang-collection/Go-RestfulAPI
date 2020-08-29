package runtime

import (
	"Go-RestfulAPI/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http"
)

/**
* @Author: super
* @Date: 2020-08-29 19:40
* @Description:
**/

func GetIP(c *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		logging.GetLogger().Error("err", zap.Error(err))
		c.String(http.StatusOK, "\n"+"error")
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				c.String(http.StatusOK, "\n"+ipnet.IP.String())
			}
		}
	}

}
