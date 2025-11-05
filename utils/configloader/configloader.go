package configloader

import (
	"github.com/spf13/viper"
	"strings"
)

func LoadConfig[T any](configPath string, envPrefix string) *T {
	v := viper.New()
	if configPath != "" {
		v.SetConfigFile(configPath)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	}
	var config T
	v.AutomaticEnv()
	v.SetEnvPrefix(envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	return &config
}