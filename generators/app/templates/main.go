package main

import (
	"flag"
	"fmt"

	"github.com/labstack/echo"

	"<%= package %>/controllers"
	"<%= package %>/models"
)

func main() {
	env := ""
	flag.StringVar(&env, "env", "", "env (local)")
	flag.Parse()
	if env == "" {
		panic("invalid env")
	}
	appCfg, err := models.GetAppCfg(env)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	e.GET("/health-check", controllers.CheckHealth)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appCfg.Port)))
}
