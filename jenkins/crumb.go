package jenkins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
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

func Request(method string, url string) (int, string, string, error) {
	var user = viper.GetString("JENKINS_USER_NAME")
	var token = viper.GetString("JENKINS_TOKEN")
	var rootURL = viper.GetString("JENKINS_ROOT_URL")

	crumbStatusCode, crumb, err := getCrumb(user, token, rootURL)
	if err != nil {
		fmt.Printf("!!! ERROR: Could not get crumb")
		return crumbStatusCode, "", "", err
	}

	// build url
	url = rootURL + url

	// log
	fmt.Printf("Method:\t\t%s\n", method)
	fmt.Printf("URL:\t\t%s\n", url)

	// create an http client
	client := &http.Client{Timeout: time.Second}

	// form the request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 400, "400", "", err
	}

	// add the crumb as a header
	req.Header.Add("Jenkins-Crumb", crumb)

	// setup auth
	req.SetBasicAuth(user, token)

	// do the request
	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, res.Status, "", err
	}

	// read the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, res.Status, "", err
	}
	res.Body.Close()

	return res.StatusCode, res.Status, string(body), err
}
