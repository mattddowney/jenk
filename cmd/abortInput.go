// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"errors"
	"fmt"
	"unicode"

	"github.com/mattddowney/jenk/jenkins"
	"github.com/spf13/cobra"
)

var jobName string

// abortInputCmd represents the abort-input command
var abortInputCmd = &cobra.Command{
	Use:   "abort-input <job_name> <build_number> <input_id>",
	Short: "Aborts a pipeline input",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 3 {
			var jobName string = args[0]
			var buildNumber string = args[1]
			var inputID string = args[2]

			// log
			fmt.Printf("Command:\tabort-input\n")
			fmt.Printf("Job Name:\t%s\n", jobName)
			fmt.Printf("Build Number:\t%s\n", buildNumber)
			fmt.Printf("Input Id:\t%s\n", inputID)

			// capitalize first letter of inputID
			inputIDRune := []rune(inputID)
			inputIDRune[0] = unicode.ToUpper(inputIDRune[0])
			inputID = string(inputIDRune)

			// build url
			url := "/job/" + jobName + "/" + buildNumber + "/input/" + inputID + "/abort"

			// issue the request
			status, _, err := jenkins.Request("POST", url)
			if err != nil {
				return err
			}

			// log status
			fmt.Printf("Status:\t\t%s\n", status)
		} else {
			return errors.New("<job_name>, <build_number>, and <input_id> required")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(abortInputCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// abortInputCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// abortInputCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
