package model

/**
* @Author: super
* @Date: 2020-08-21 20:10
* @Description:
**/

type RedisConfig struct {
	MaxIdle   int    `mapstructure:"maxIdle"`
	MaxActive int    `mapstructure:"max_activate"`
	Url       string `mapstructure:"host"`
}
