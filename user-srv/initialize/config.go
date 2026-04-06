package initialize

import (
	"fmt"

	"github.com/1348453525/user-redeem-code-grpc/user-srv/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitConfig(cfgFile string) {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))

	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file is changed: ", in.Name)
		if err := viper.Unmarshal(&global.Config); err != nil {
			panic(fmt.Errorf("failed to reload config: %w", err))
		}
	})
}
