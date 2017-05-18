package jenkins

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/spf13/viper"
)

func Request(method string, url string, body *url.Values) (int, string, string, error) {
	var user = viper.GetString("JENKINS_USER_NAME")
	var token = viper.GetString("JENKINS_TOKEN")
	var rootURL = viper.GetString("JENKINS_ROOT_URL")

	// get the crumb
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
	fmt.Printf("Request Body:\t%s\n", body.Encode())

	// create an http client
	client := &http.Client{Timeout: time.Second}

	// convert body into a buffer
	bodyBuff := bytes.NewBufferString(body.Encode())

	// form the request
	req, err := http.NewRequest(method, url, bodyBuff)
	if err != nil {
		return 400, "400", "", err
	}

	// add headers
	req.Header.Add("Jenkins-Crumb", crumb)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// setup auth
	req.SetBasicAuth(user, token)

	// do the request
	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, res.Status, "", err
	}

	// read the body
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, res.Status, "", err
	}
	res.Body.Close()

	return res.StatusCode, res.Status, string(resBody), err
}
