package main

import (
	"encoding/json"
	"io/ioutil"
)

type AppConfig struct {
	Port                string `json:"port"`
	ApiEp               string `json:"apiep"`
	GSApi1              string `json:"gsapi1"`
	GSApi2              string `json:"gsapi2"`
	HackathonMemePicURL string `json:"hackathonMemePicURL"`
}

var (
	AppConf *AppConfig
)

func InitAppConfig(file string) error {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(raw, &AppConf); err != nil {
		return err
	}
	return nil
}

func GetAppConfig() *AppConfig {
	return AppConf
}
