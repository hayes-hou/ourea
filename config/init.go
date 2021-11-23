// Package config Copyright (c) 2021 ~ 2022 Ourea Team
// Implementation of DDD framework with golang
package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	Name string
}

func (conf *Conf) getConfig() error {
	viper.AddConfigPath("./")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func InitConfig(cfg string) error {
	conf := Conf{
		Name: cfg,
	}
	if err := conf.getConfig(); err != nil {
		return err
	}
	return nil
}

func GetKeyByConf(key, t string) interface{} {
	if key == "" {
		return ""
	}
	switch t {
	case "str":
		return viper.GetString(key)
	case "bool":
		return viper.GetBool(key)
	case "int":
		return viper.GetInt(key)
	case "uint64":
		return viper.GetUint64(key)
	case "float64":
		return viper.GetFloat64(key)
	case "duration":
		return viper.GetDuration(key)
	}
	return ""
}
