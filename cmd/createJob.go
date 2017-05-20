// Copyright © 2017 Matthew Downey <mattddowney@gmail.com>
//

package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"

	"github.com/mattddowney/jenk/jenkins"
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

			type project struct {
				ProjectURL string
				GitURL     string
			}

			const newJobTemplate string = `
				<flow-definition plugin="workflow-job@2.9">
					<keepDependencies>false</keepDependencies>
					<properties>
						<com.coravy.hudson.plugins.github.GithubProjectProperty plugin="github@1.25.1">
							<projectUrl>{{ .ProjectURL }}</projectUrl>
						</com.coravy.hudson.plugins.github.GithubProjectProperty>
					</properties>
					<definition class="org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition" plugin="workflow-cps@2.25">
						<scm class="hudson.plugins.git.GitSCM" plugin="git@3.0.1">
							<configVersion>2</configVersion>
							<userRemoteConfigs>
								<hudson.plugins.git.UserRemoteConfig>
									<url>{{ .GitURL }}</url>
								</hudson.plugins.git.UserRemoteConfig>
							</userRemoteConfigs>
							<branches>
								<hudson.plugins.git.BranchSpec>
									<name>*/master</name>
								</hudson.plugins.git.BranchSpec>
							</branches>
							<doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>
							<submoduleCfg class="list"/>
							<extensions>
								<hudson.plugins.git.extensions.impl.SubmoduleOption>
									<disableSubmodules>false</disableSubmodules>
									<recursiveSubmodules>true</recursiveSubmodules>
									<trackingSubmodules>false</trackingSubmodules>
									<parentCredentials>false</parentCredentials>
								</hudson.plugins.git.extensions.impl.SubmoduleOption>
							</extensions>
						</scm>
						<scriptPath>Jenkinsfile</scriptPath>
					</definition>
				</flow-definition>`

			if gitURL == "" {
				gitURL = projectURL + ".git"
			}

			// log
			fmt.Printf("Command:\tcreate-job\n")
			fmt.Printf("Job Name:\t%s\n", jobName)
			fmt.Printf("Project URL:\t%s\n", projectURL)
			fmt.Printf("Git URL:\t%s\n", gitURL)

			// build url
			reqURL := "/createItem?name=" + jobName

			// form body from
			tmpl, err := template.New("body").Parse(newJobTemplate)
			if err != nil {
				panic(err)
			}
			projectStruct := project{projectURL, gitURL}
			var body bytes.Buffer
			err = tmpl.Execute(&body, projectStruct)
			if err != nil {
				panic(err)
			}

			// issue the request
			statusCode, status, _, err := jenkins.Request("POST", reqURL, "application/xml", &body)
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
