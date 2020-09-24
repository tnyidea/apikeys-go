package apikeys

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func (m *ApiKeyMap) ValidateAPIKeyHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := mux.Vars(r)["key"]

		path, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Error", 500)
			return
		}
		path = strings.ReplaceAll(path, "/{key}", "")

		err = m.ValidateApiKeyDefault(key, path)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		handler(w, r)
	}
}
