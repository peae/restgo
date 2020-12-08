package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Addr		string		`yaml:"addr"`
	DSN			string		`yaml:"dsn"`
	MaxIdleConn	int			`yaml:"max_idle_conn"`
}

var config *Config

func Load(path string) error {
	fmt.Println(path)
	fmt.Println("=====")
	result, err := ioutil.ReadFile(path)
	fmt.Println(result)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(result, &config)
}

func Get() *Config {
	return config
}
