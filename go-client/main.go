package main

import (
	"flag"
	"fmt"
	"go-client/config"
	"net/http"

	"github.com/spf13/viper"
)

var appName = "spring-config-server"
var amqp_server_url = "amqp://guest:guest@localhost:5672/"
var config_event_bus = "springCloudBus"

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

func handler(resp http.ResponseWriter, req *http.Request) {

	fmt.Println("Request Recieved")
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	resp.Write([]byte(viper.GetString("server.message")))

}

func main() {
	fmt.Printf("Starting %v\n", appName)

	config.LoadConfigurationFromBranch(
		viper.GetString("configServerUrl"),
		appName,
		viper.GetString("profile"),
		viper.GetString("configBranch"))

	go config.StartListener(appName, amqp_server_url, config_event_bus)
	http.HandleFunc("/server/message", handler)

	http.ListenAndServe(":8066", nil)

	// service.StartWebServer(viper.GetString("server_port"))
}
