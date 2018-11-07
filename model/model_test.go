package model

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewMagazine(t *testing.T) {

	// var json := "{\"magazine_id\":\"b09c2d83-498a-4f8a-a614-358d188cea0a\", \"author\": \"Jack Wright\", \"title "
	uuidString := "b09c2d83-498a-4f8a-a614-358d188cea0a"
	myUUID, err := uuid.Parse(uuidString)
	title := "National Geographic"
	issue := "October 2018"

	if err == nil {
		magazine := NewMagazine(myUUID, title, issue)

		if magazine.Title != title {
			t.Errorf("Expected Title to be %s not %s", title, magazine.Title)
		}

		if magazine.Issue != issue {
			t.Errorf("Expected Issue to be %s not %s", issue, magazine.Issue)
		}

		if magazine.MagazineID != myUUID {
			t.Errorf("Expected MagazineID to be %s not %s", myUUID, magazine.MagazineID)
		}

	} else {
		t.Errorf("Unable to create a uuid: %s", err)
	}

}

func TestJSON(t *testing.T) {
	uuidString := "b09c2d83-498a-4f8a-a614-358d188cea0a"
	myUUID, err := uuid.Parse(uuidString)
	title := "National Geographic"
	issue := "October 2018"
	expectedJSON := "{\"MagazineID\":\"b09c2d83-498a-4f8a-a614-358d188cea0a\",\"Title\":\"National Geographic\",\"Issue\":\"October 2018\"}"

	if err == nil {
		magazine := NewMagazine(myUUID, title, issue)
		jsonBytes, err := magazine.JSON()

		if err == nil {

			if string(jsonBytes) != expectedJSON {
				t.Errorf("Expected json of %s received %s", expectedJSON, jsonBytes)
			}

		} else {
			t.Errorf("Unable to marshall Magazine to json err: %s", err)
		}

	} else {
		t.Errorf("Unable to create a uuid: %s", err)
	}

}

func TestNewMagazineFromJSON(t *testing.T) {
	json := "{\"MagazineID\":\"b09c2d83-498a-4f8a-a614-358d188cea0a\",\"Title\":\"National Geographic\",\"Issue\":\"October 2018\"}"
	magazine, err := NewMagazineFromJSON([]byte(json))
	if err == nil {
		uuidString := "b09c2d83-498a-4f8a-a614-358d188cea0a"
		myUUID, _ := uuid.Parse(uuidString)
		if magazine.MagazineID != myUUID {
			t.Errorf("Expected MagazineID of %s, received %s", myUUID, magazine.MagazineID)
		}

		title := "National Geographic"
		if title != magazine.Title {
			t.Errorf("Expected Title of %s, received %s", title, magazine.Title)
		}

		issue := "October 2018"
		if issue != magazine.Issue {
			t.Errorf("Expected Issue of %s, received %s", issue, magazine.Issue)
		}

	} else {
		t.Errorf("Unable to unmarshal: %s", magazine)
	}
}
