package service

import (
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gziphandler"
	"github.com/sirupsen/logrus"
)

type (
	// MagazineService handles CRUD magazine actions
	MagazineService struct {
	}

	// Config for MagazineService
	Config struct {
		Server *server.Config
	}
)

// NewMagazineService creates a new instance of MagazineService with the specified config
func NewMagazineService(cfg *Config) *MagazineService {
	return &MagazineService{}
}

// Prefix defines the url prefix this service is mapped to (From server.JsonService)
func (s *MagazineService) Prefix() string {
	return "/magazine"
}

// Middleware defines middleware implementation (From server.JsonService)
func (s *MagazineService) Middleware(h http.Handler) http.Handler {
	return gziphandler.GzipHandler(h)
}

// JSONMiddleware configures this service's json handling
func (s *MagazineService) JSONMiddleware(j server.JSONEndpoint) server.JSONEndpoint {
	return func(r *http.Request) (int, interface{}, error) {

		status, res, err := j(r)
		if err != nil {
			server.LogWithFields(r).WithFields(logrus.Fields{
				"error": err,
			}).Error("problems with serving request")
			return http.StatusServiceUnavailable, nil, &jsonErr{"sorry, this service is unavailable"}
		}

		server.LogWithFields(r).Info("success!")
		return status, res, nil
	}

}

// func (s *MagazineService) JSONEndpoints() map[string]map[string]server.JSONEndpoint {
// 	"/{magazine_id}": map[string]server.JSONEndpoint {
// 		"GET": s
// 	}

// }

type jsonErr struct {
	Err string `json:"error"`
}

func (e *jsonErr) Error() string {
	return e.Err
}
