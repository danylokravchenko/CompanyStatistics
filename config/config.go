package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)


// Init viper to read config file
func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml") // config file extension
	viper.AddConfigPath("config") // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	viper.WatchConfig() // watch for changes in config
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		fmt.Println("Config file changed:", e.Name)
	})

}


// Get config setting
func GetSetting(key string) interface{}{
	return viper.Get(key)
}
