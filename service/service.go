package service

import (
	"io/ioutil"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/NYTimes/gizmo/web"
	"github.com/NYTimes/gziphandler"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"github.com/ayax79/go-magazines/dao"
	"github.com/ayax79/go-magazines/model"
)

type (
	// MagazineService handles CRUD magazine actions
	MagazineService struct {
		magazineDAO *dao.RedisMagazineDAO
	}

	// Config for MagazineService
	Config struct {
		Server *server.Config
	}
)

// NewMagazineService creates a new instance of MagazineService with the specified config
func NewMagazineService(cfg *Config, redisCfg *dao.RedisConfig) (*MagazineService, error) {
	d, err := dao.NewRedisMagazineDAO(redisCfg)
	if d != nil {
		return &MagazineService{
			magazineDAO: d,
		}, err
	} else {
		return nil, err
	}
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

// JSONEndpoints provides url mappings
func (s *MagazineService) JSONEndpoints() map[string]map[string]server.JSONEndpoint {
	return map[string]map[string]server.JSONEndpoint{
		"/{magazine_id}": map[string]server.JSONEndpoint{
			"GET": s.getMagazine,
		},
		"/": map[string]server.JSONEndpoint{
			"POST": s.postMagazine,
		},
	}
}

func (s *MagazineService) getMagazine(r *http.Request) (int, interface{}, error) {
	uuidString := web.Vars(r)["magazine_id"]
	magazineID, err := uuid.Parse(uuidString)
	if err != nil {
		magazine, err := s.magazineDAO.Get(magazineID)
		if err != nil {
			json, err := magazine.JSON()
			if err != nil {
				return http.StatusOK, json, err
			}
			return http.StatusInternalServerError, nil, err
		}
		return http.StatusNotFound, nil, err
	}
	return http.StatusBadRequest, nil, err
}

func (s *MagazineService) postMagazine(r *http.Request) (int, interface{}, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	magazine, err := model.NewMagazineFromJSON(b)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err2 := s.magazineDAO.Put(&magazine)
	if err2 != nil {
		return http.StatusInternalServerError, nil, err2
	}
	return http.StatusOK, nil, err2
}

type jsonErr struct {
	Err string `json:"error"`
}

func (e *jsonErr) Error() string {
	return e.Err
}
