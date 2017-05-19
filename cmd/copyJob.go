// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/mattddowney/jenk/jenkins"
	"github.com/spf13/cobra"
)

// copyJobCmd represents the copy-job command
var copyJobCmd = &cobra.Command{
	Use:   "copy-job <from_job_name> <new_job_name>",
	Short: "Copy a Jenkins job",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			var fromJobName = args[0]
			var newJobName = args[1]

			// log
			fmt.Printf("Command:\tcopy-job\n")
			fmt.Printf("From Job Name:\t%s\n", fromJobName)
			fmt.Printf("New Job Name:\t%s\n", newJobName)

			// build url
			reqURL := "/createItem?name=" + newJobName + "&mode=copy&from=" + fromJobName

			// create body
			body := url.Values{}
			body.Set("json", "{\"parameter\": []}")

			// issue the request
			statusCode, status, _, err := jenkins.Request("POST", reqURL, &body)
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
			return errors.New("<from_job_name>, <new_job_name> required")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(copyJobCmd)
}
