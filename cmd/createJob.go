// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var gitURL string

// createJobCmd represents the create-job command
var createJobCmd = &cobra.Command{
	Use:   "create-job <job_name> <project_url>",
	Short: "A brief description of your command",

	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("create-job called")
		if len(args) == 2 {
			var jobName = args[0]
			var projectURL = args[1]

			if gitURL == "" {
				gitURL = projectURL + ".git"
			}

			// log
			fmt.Printf("Command:\tcreate-job\n")
			fmt.Printf("Job Name:\t%s\n", jobName)
			fmt.Printf("Project URL:\t%s\n", projectURL)
			fmt.Printf("Git URL:\t%s\n", gitURL)
		} else {
			return errors.New("<job_name> and <project_url> required")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(createJobCmd)

	// flags
	createJobCmd.Flags().StringVarP(&gitURL, "git-url", "g", "", "git URL (defaults to <project_url>.git)")
}
