package security

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AuthorizeTransaction() bool {
	resp, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		return false
	}

	var body struct {
		Authorization bool `json:"authorization"`
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return false
	}

	return body.Authorization
}
