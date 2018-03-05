package utils

import (
	"encoding/json"
	"net/http"
	"riggedstars/app/models"
)

func ParseBody(req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t models.User
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()

}
