// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "Print the current environment",
	Long: `Print the current environment.
These are used to interact with your Jenkins installation.
Environment variables override the default config at $HOME/.jenk.yaml`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("JENKINS_ROOT_URL=%s\n", viper.GetString("jenkins_root_url"))
		fmt.Printf("JENKINS_TOKEN=%s\n", viper.GetString("jenkins_token"))
		fmt.Printf("JENKINS_USER_NAME=%s\n", viper.GetString("jenkins_user_name"))
	},
}

func init() {
	RootCmd.AddCommand(envCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// envCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// envCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
