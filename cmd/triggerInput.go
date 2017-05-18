// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// triggerInputCmd represents the trigger-input command
var triggerInputCmd = &cobra.Command{
	Use:   "trigger-input",
	Short: "Triggers a pipeline input",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("triggerInput called")
	},
}

func init() {
	RootCmd.AddCommand(triggerInputCmd)
}
