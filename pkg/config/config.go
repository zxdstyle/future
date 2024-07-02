package config

import (
	"github.com/spf13/viper"
	"log"
)

func Load() {
	//remote.SetAppID(os.Getenv("APOLLO_APP_ID"))
	//remote.SetConfigType("yaml")
	//remote.SetAgolloOptions(agollo.DefaultNamespace("application.yaml"), agollo.AutoFetchOnCacheMiss())
	viper.SetConfigFile("yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	//err := viper.AddRemoteProvider("apollo", os.Getenv("APOLLO_CONFIG_SERVER"), os.Getenv("APOLLO_NAMESPACES"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err = viper.ReadRemoteConfig(); err != nil {
	//	log.Fatal(err)
	//}
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
