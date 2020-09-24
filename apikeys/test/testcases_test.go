package test

import (
	"github.com/tnyidea/apikeys-go/apikeys"
	"log"
	"testing"
)

var testApiKeyMap apikeys.ApiKeyMap

func TestGetApiKeyMapFromFile(t *testing.T) {
	apiKeyMap, err := apikeys.GetApiKeyMapFromFile("sample-api-keys.json")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	testApiKeyMap = apiKeyMap
}

func TestValidateValidApiKeyValidUri(t *testing.T) {
	key := "Kn0NKHABiruvX3FM7s0J"
	uri := "/sample/uri/two"

	err := testApiKeyMap.ValidateApiKeyDefault(key, uri)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
}

func TestValidateInvalidApiKeyValidUri(t *testing.T) {
	key := "badbadbadbadbadbad"
	uri := "/sample/uri/two"

	err := testApiKeyMap.ValidateApiKeyDefault(key, uri)
	if err == nil {
		log.Println("ERROR: Invalid key validated for valid uri")
		t.FailNow()
	}
}

func TestValidateValidApiKeyInvalidUri(t *testing.T) {
	key := "Kn0NKHABiruvX3FM7s0J"
	uri := "/bad/bad/uri"

	err := testApiKeyMap.ValidateApiKeyDefault(key, uri)
	if err == nil {
		log.Println("ERROR: Valid key validated for invalid uri")
		t.FailNow()
	}
}

func TestValidateInvalidApiKeyInvalidUri(t *testing.T) {
	key := "badbadbadbadbadbad"
	uri := "/bad/bad/uri"

	err := testApiKeyMap.ValidateApiKeyDefault(key, uri)
	if err == nil {
		log.Println("ERROR: Invalid key validated for invalid uri")
		t.FailNow()
	}
}

// TODO Create a test case for the HTTP Handler
