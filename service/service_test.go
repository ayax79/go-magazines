package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/NYTimes/gizmo/server"
	"github.com/ayax79/go-magazines/dao"
	"github.com/ayax79/go-magazines/model"
	"github.com/google/uuid"
	"github.com/ory/dockertest"
)

var srvr *server.SimpleServer

func TestMain(m *testing.M) {

	//configure docker
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("redis", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		// configure the test server
		redisHost := fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp"))
		redisConfig := dao.NewRedisConfig(redisHost, "", 0)

		srvr = server.NewSimpleServer(nil)

		svc, err := NewMagazineService(nil, redisConfig)
		if err != nil {
			log.Fatalf("Could not create service %s", err)
			return err
		}
		srvr.Register(svc)

		return nil
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}

func TestCRUD(t *testing.T) {
	uuid, _ := uuid.NewUUID()
	title := "Guitar World"
	issue := "June 2018"
	magazine := model.NewMagazine(uuid, title, issue)
	jsonBytes, _ := magazine.JSON()
	reader := bytes.NewReader(jsonBytes)
	r, _ := http.NewRequest("POST", "/magazine", reader)
	w := httptest.NewRecorder()
	srvr.ServeHTTP(w, r)

	if w.Code != 202 {
		t.Errorf("Expected status 202, received %#v", w.Code)
		t.FailNow()
	}

	r2, _ := http.NewRequest("GET", fmt.Sprintf("/magazine/%s", uuid.String()), nil)
	w2 := httptest.NewRecorder()
	srvr.ServeHTTP(w2, r2)

	if w2.Code != 200 {
		t.Errorf("Expected status 200, received %#v", w2.Code)
		t.FailNow()
	}

	magazineResponse, err := model.NewMagazineFromJSON(w2.Body.Bytes())

	if err != nil {
		t.Errorf("Received an err while parsing json response: %s", err)
		t.FailNow()
	}

	if magazine.MagazineID != magazineResponse.MagazineID {
		t.Errorf("Expected MagazineID of %s, received %s", magazine.MagazineID, magazineResponse.MagazineID)
	}

	if magazine.Title != magazineResponse.Title {
		t.Errorf("Expected Title of %s, received %s", magazine.Title, magazineResponse.Title)
	}

	if magazine.Issue != magazineResponse.Issue {
		t.Errorf("Expected Issue of %s, received %s", magazine.Issue, magazineResponse.Issue)
	}

}
