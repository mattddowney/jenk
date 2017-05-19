// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// copyJobCmd represents the copy-job command
var copyJobCmd = &cobra.Command{
	Use:   "copy-job",
	Short: "Copy a Jenkins job",

	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 2 {
			var fromJobName = args[0]
			var newJobName = args[1]

			// log
			fmt.Printf("Command:\tcopy-job\n")
			fmt.Printf("From Job Name:\t%s\n", fromJobName)
			fmt.Printf("New Job Name:\t%s\n", newJobName)
		} else {
			return errors.New("<from_job_name>, <new_job_name> required")
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(copyJobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// copyJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// copyJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
