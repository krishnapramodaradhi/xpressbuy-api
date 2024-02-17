package entity

import "os"

type HealthCheck struct {
	Id          int    `json:"id"`
	Environment string `json:"environment"`
	Status      string `json:"status"`
}

func NewHealthCheck() HealthCheck {
	return HealthCheck{Id: os.Getpid(), Environment: os.Getenv("APP_ENV"), Status: "healthy"}
}
