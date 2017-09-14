package models

import (
	"errors"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type AppCfg struct {
	Port int
}

type AppCfgs struct {
	Local AppCfg
}

var appCfgs = AppCfgs{}

func (cfgs *AppCfgs) Read(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, cfgs)
	if err != nil {
		return err
	}
	return nil
}

func GetAppCfg(env string) (AppCfg, error) {
	if appCfgs.Local.Port == 0 {
		err := appCfgs.Read("configs/app.yml")
		if err != nil {
			return AppCfg{}, err
		}
	}
	if env == "local" {
		return appCfgs.Local, nil
	}
	return appCfgs.Local, errors.New(fmt.Sprintf("invalid env: %s", env))
}
