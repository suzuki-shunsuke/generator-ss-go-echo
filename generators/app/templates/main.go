package main

import (
	"fmt"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"

	"<%= package %>/controllers"
)

var (
	envs    []string = []string{}
	intEnvs []string = []string{"app_port"}
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	bindEnvs()
	setFlags()
}

func main() {
	if err := validateFlag(); err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.GET("/healthz", controllers.CheckHealth)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", viper.GetInt("app_port"))))
}

func bindEnvs() {
	for _, e := range intEnvs {
		viper.BindEnv(e)
	}
	for _, e := range envs {
		viper.BindEnv(e)
	}
}

func setFlags() {
	flag.Int("port", 1323, "")
	flag.Parse()
	viper.BindPFlag("app_port", flag.Lookup("port"))
}

func validateFlag() error {
	for _, e := range intEnvs {
		if viper.GetInt(e) == 0 {
			return fmt.Errorf("[FAIRURE] %s is required", e)
		}
	}
	for _, e := range envs {
		if viper.GetString(e) == "" {
			return fmt.Errorf("[FAIRURE] %s is required", e)
		}
	}
	return nil
}
