// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"errors"
	"fmt"
	"net/url"
	"unicode"

	"github.com/mattddowney/jenk/jenkins"
	"github.com/spf13/cobra"
)

// triggerInputCmd represents the trigger-input command
var triggerInputCmd = &cobra.Command{
	Use:   "trigger-input",
	Short: "Triggers a pipeline input",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 3 {
			var jobName string = args[0]
			var buildNumber string = args[1]
			var inputID string = args[2]

			// log
			fmt.Printf("Command:\ttrigger-input\n")
			fmt.Printf("Job Name:\t%s\n", jobName)
			fmt.Printf("Build Number:\t%s\n", buildNumber)
			fmt.Printf("Input Id:\t%s\n", inputID)

			// capitalize first letter of inputID
			inputIDRune := []rune(inputID)
			inputIDRune[0] = unicode.ToUpper(inputIDRune[0])
			inputID = string(inputIDRune)

			// build url
			reqUrl := "/job/" + jobName + "/" + buildNumber + "/wfapi/inputSubmit?inputId=" + inputID

			body := url.Values{}
			body.Set("json", "{\"parameter\": []}")

			// issue the request
			statusCode, status, _, err := jenkins.Request("POST", reqUrl, &body)
			if err != nil {
				return err
			}

			// log status
			fmt.Printf("StatusCode:\t%d\n", statusCode)
			fmt.Printf("Status:\t\t%s\n", status)

			if statusCode != 200 {
				return errors.New("Request failed")
			}
		} else {
			return errors.New("<job_name>, <build_number>, and <input_id> required")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(triggerInputCmd)
}
