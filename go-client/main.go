package main

import (
	"flag"
	"fmt"
	"go-client/config"

	"github.com/spf13/viper"
)

var appName = "spring-config-server"

func init() {
	profile := flag.String("profile", "development", "Environment profile, something similar to spring profiles")
	configServerUrl := flag.String("configServerUrl", "http://localhost:8888", "Address to config server")
	configBranch := flag.String("configBranch", "main", "git branch to fetch configuration from")

	flag.Parse()

	fmt.Println("Specified configBranch is " + *configBranch)

	viper.Set("profile", *profile)
	viper.Set("configServerUrl", *configServerUrl)
	viper.Set("configBranch", *configBranch)
}

func main() {
	fmt.Printf("Starting %v\n", appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	// go config.StartListener(appName, viper.GetString("amqp_server_url"), viper.GetString("config_event_bus"))
	// service.StartWebServer(viper.GetString("server_port"))
}
