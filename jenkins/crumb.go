package jenkins

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func getCrumb(user string, token string, rootURL string) (int, string, error) {
	// full url to the crumb issuer on the Jenkins server
	var url = rootURL + "/crumbIssuer/api/json"

	// create an http client
	client := &http.Client{Timeout: time.Second}

	// form the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 400, "", err
	}

	// setup auth
	req.SetBasicAuth(user, token)

	// do the request
	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, "", err
	}

	// read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 500, "", err
	}
	res.Body.Close()

	// parse json
	var j interface{}
	err = json.Unmarshal(body, &j)
	if err != nil {
		return 500, "", err
	}
	mapped := j.(map[string]interface{})

	// return the crumb
	return 0, mapped["crumb"].(string), nil
}
