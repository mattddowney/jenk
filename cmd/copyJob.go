// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// copyJobCmd represents the copy-job command
var copyJobCmd = &cobra.Command{
	Use:   "copy-job",
	Short: "Copy a Jenkins job",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("copyJob called")
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
