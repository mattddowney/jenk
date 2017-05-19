// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createJobCmd represents the create-job command
var createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create-job called")
	},
}

func init() {
	RootCmd.AddCommand(createJobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
