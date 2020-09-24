package apikeys

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"time"
)

type ApiKey struct {
	Id            int       `json:"id"`
	UserId        string    `json:"userId"`
	Key           string    `json:"key"`
	Expiration    time.Time `json:"expiration"`
	Permissions   []string  `json:"permissions"`
	PermissionMap map[string]bool
}

type ApiKeyMap struct {
	keyMap map[string]ApiKey
}

func (a *ApiKey) Bytes() []byte {
	byteValue, _ := json.Marshal(a)
	return byteValue
}

func (a *ApiKey) String() string {
	byteValue, _ := json.MarshalIndent(a, "", "    ")
	return string(byteValue)
}

func GetApiKeyMapFromFile(filename string) (ApiKeyMap, error) {
	var apiKeys []ApiKey
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return ApiKeyMap{}, err
	}
	err = json.Unmarshal(b, &apiKeys)
	if err != nil {
		log.Println("initialize API keys:", err)
		return ApiKeyMap{}, err
	}

	apiKeyMap := make(map[string]ApiKey)
	for _, apiKey := range apiKeys {
		// Build permission map
		permissionMap := make(map[string]bool)
		for _, permission := range apiKey.Permissions {
			permissionMap[permission] = true
		}
		apiKey.PermissionMap = permissionMap

		// Lookup key information by key or id
		apiKeyMap[apiKey.UserId] = apiKey
		apiKeyMap[apiKey.Key] = apiKey
	}

	return ApiKeyMap{
		keyMap: apiKeyMap,
	}, nil
}

func (m *ApiKeyMap) ValidateApiKeyDefault(key string, uri string) error {
	errMessage := "permission denied for URI: " + uri
	if _, defined := m.keyMap[key]; !defined {
		log.Println(errMessage + ": key does not exist")
		return errors.New(errMessage)
	}

	errMessage += " for user: " + m.keyMap[key].UserId
	if !m.keyMap[key].PermissionMap[uri] {
		log.Println(errMessage + ": permission denied")
		return errors.New(errMessage)
	}
	if m.keyMap[key].Expiration.Before(time.Now()) {
		log.Println(errMessage + ": key expired")
		return errors.New(errMessage)
	}

	return nil
}
