package services

import (
	"api-gateway/config"
)

type ServiceRegistry struct {
	Config config.Config
}

func NewServiceRegistry(config config.Config) *ServiceRegistry {
	return &ServiceRegistry{Config: config}
}

func (sr *ServiceRegistry) GetServiceURL(serviceName string) string {
	switch serviceName {
	case "user":
		return sr.Config.UserServiceURL
	case "product":
		return sr.Config.ProductServiceURL
	case "order":
		return sr.Config.OrderServiceURL
	default:
		return ""
	}
}
