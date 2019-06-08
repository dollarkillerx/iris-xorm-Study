/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午5:52
* */
package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DriverName string `json:"driver_name"`
	Dsn string `json:"dsn"`
	Debug bool `json:"debug"`
}

var (
	Configs *Config
)

func init() {
	file, e := os.Open("./config.json")
	if e != nil {
		panic(e.Error())
	}

	Configs = &Config{}
	decoder := json.NewDecoder(file)
	e = decoder.Decode(Configs)
	if e != nil {
		panic(e.Error())
	}
}

