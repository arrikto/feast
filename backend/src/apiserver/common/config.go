package common

import (
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

func GetStringConfig(configName string) string {
	if !viper.IsSet(configName) {
		glog.Fatalf("Please specify flag %s", configName)
	}

	return viper.GetString(configName)
}

func GetStringConfigWithDefault(configName, value string) string {
	if !viper.IsSet(configName) {
		return value
	}

	return viper.GetString(configName)
}

func GetMapConfig(configName string) map[string]string {
	if !viper.IsSet(configName) {
		glog.Infof("Config %s not specified, skipping", configName)
		return nil
	}

	return viper.GetStringMapString(configName)
}

func GetBoolConfigWithDefault(configName string, value bool) bool {
	if !viper.IsSet(configName) {
		return value
	}

	value, err := strconv.ParseBool(viper.GetString(configName))
	if err != nil {
		glog.Fatalf("Failed converting string to bool %s", viper.GetString(configName))
	}

	return value
}

func GetFloat64ConfigWithDefault(configName string, value float64) float64 {
	if !viper.IsSet(configName) {
		return value
	}

	return viper.GetFloat64(configName)
}

func GetIntConfigWithDefault(configName string, value int) int {
	if !viper.IsSet(configName) {
		return value
	}

	return viper.GetInt(configName)
}

func GetDurationConfig(configName string) time.Duration {
	if !viper.IsSet(configName) {
		glog.Fatalf("Please specify flag %s", configName)
	}

	return viper.GetDuration(configName)
}
