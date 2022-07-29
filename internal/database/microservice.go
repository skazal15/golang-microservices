package database

import (
	"net/http"
	"time"
)

const getUserUrl = "internal/users"

type MicroserviceRepo struct {
	hostname string
	username string
	password string
	client   *http.Client
}

func Microservice() *MicroserviceRepo {
	cfg, _ := ConfigYaml()
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	return &MicroserviceRepo{
		hostname: cfg.APIhost,
		username: cfg.APIuser,
		password: cfg.APIpass,
		client:   &client,
	}
}
