package dao

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ayax79/go-magazines/model"
	"github.com/google/uuid"
	"github.com/ory/dockertest"
)

var dao *RedisMagazineDAO

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("redis", "latest", nil)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		host := fmt.Sprintf("localhost:%s", resource.GetPort("6379/tcp"))
		config := NewRedisConfig(host, "", 0)
		d, err := NewRedisMagazineDAO(config)
		dao = d
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	code := m.Run()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestCrud(t *testing.T) {
	magazineID, _ := uuid.NewUUID()
	title := "New York Times Magazine"
	issue := "November 2018"
	magazine := model.NewMagazine(magazineID, title, issue)
	err := dao.Put(magazine)
	if err != nil {
		t.Errorf("Error putting %s into redis: %s", magazine, err)
	} else {
		magazine, err := dao.Get(magazineID)

		if magazine != nil {

			if magazine.MagazineID != magazineID {
				t.Errorf("Expected MagazineID %s but received %s", magazineID, magazine.MagazineID)
			}

			if magazine.Title != title {
				t.Errorf("Expected Title %s but received %s", title, magazine.Title)
			}

			if magazine.Issue != issue {
				t.Errorf("Expected Issue %s but received %s", issue, magazine.Issue)
			}

		} else if err != nil {
			t.Errorf("Error retrieving magazine with id %s from redis: %s", magazineID, err)
		} else {
			t.Errorf("Entry with id %s not found in redis", magazineID)
		}

	}

}
