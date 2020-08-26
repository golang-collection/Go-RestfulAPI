package model

import "time"

/**
* @Author: super
* @Date: 2020-08-21 21:19
* @Description:
**/

type ServerConfig struct {
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
