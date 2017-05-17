package jenkins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

func GetCrumb() error {
	var user = viper.GetString("JENKINS_USER_NAME")
	var token = viper.GetString("JENKINS_TOKEN")
	var rootURL = viper.GetString("JENKINS_ROOT_URL")

	// full url to the crumb issuer on the Jenkins server
	var url = rootURL + "/crumbIssuer/api/json"

	client := &http.Client{Timeout: time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// setup auth
	req.SetBasicAuth(user, token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	res.Body.Close()

	fmt.Printf("%s", body)

	return nil
}
