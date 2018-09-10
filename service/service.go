package service

import (
	"github.com/NYTimes/gizmo/server"
)

type (
	JSONService interface {
	}

	Config struct {
		Server           *server.Config
		MostPopularToken string
		SemanticToken    string
	}
)

func NewJSONService(cfg *Config) *JSONService {
	return &JSONService{}
}
