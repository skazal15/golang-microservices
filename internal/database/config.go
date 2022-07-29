package database

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	Host              string `yaml:"host"`
	Port              string `yaml:"port"`
	Dbname            string `yaml:"dbname"`
	MaxIdleUser       int    `yaml:"max_idle_user"`
	MaxUserConnection int    `yaml:"max_connection"`
	APIhost           string `yaml:"API_HOST"`
	APIuser           string `yaml:"API_USER"`
	APIport           string `yaml:"API_PORT"`
	APIpass           string `yaml:"API_PASSWORD"`
}

type configurl struct {
	Urllogin           string `yaml:"login"`
	Urlregister        string `yaml:"register"`
	UrlGetExerciseById string `yaml:"get_exercise_by_id"`
	UrlCreateExercise  string `yaml:"create_exercise"`
	UrlCreateQuestion  string `yaml:"create_question"`
	UrlCreateAnswer    string `yaml:"create_answer"`
	UrlCalculateScore  string `yaml:"calcute_user_score"`
}

func ConfigYaml() (conf *config, err error) {
	data, err := ioutil.ReadFile("./internal/database/config.yaml")
	if err != nil {
		fmt.Println("file tidak ketemu")
	}
	conf = &config{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
	return
}

func ConfigUrl() (conf *configurl, err error) {
	data, err := ioutil.ReadFile("./internal/database/config_url.yaml")
	if err != nil {
		fmt.Println("file tidak ketemu")
	}
	conf = &configurl{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
	return
}
