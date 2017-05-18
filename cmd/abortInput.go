// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"errors"
	"fmt"

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

			fmt.Printf("Job Name: %s\n", jobName)
			fmt.Printf("Build Number: %s\n", buildNumber)
			fmt.Printf("Input Id: %s\n", inputID)
		} else {
			return errors.New("<job_name>, <build_number>, and <input_id> required")
		}

		var crumb, err = jenkins.GetCrumb()
		if err != nil {
			return err
		}

		fmt.Printf("Crumb: %s\n", crumb)

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
