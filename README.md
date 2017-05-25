jenk
====

Interact with [Jenkins](https://jenkins.io/) via the command line. Written in Go. 
Rewrite of [jenk-bash](https://github.com/mattddowney/jenk-bash).

Copy jobs, create pipelines, and trigger pipeline inputs.

Setup:
------

Set the following environment variables:

* JENKINS_ROOT_URL - The url of the Jenkins server (IE: http://localhost:8080)
* JENKINS_USER_NAME - Jenkins user with API access
* JENKINS_TOKEN - Access token obtained from Jenkins (http://localhost:8080/me/configure)

*OR*

Create a configuration file named `.jenk.yaml` in your home directory containing same variables, 
with the following layout:

```yaml
JENKINS_ROOT_URL: http://localhost:8080
JENKINS_USER_NAME: username
JENKINS_TOKEN: token_from_jenkins
```

Usage:
------

```
$ ./jenk
A tool for interacting with the Jenkins API.

Usage:
  jenk [command]

Available Commands:
  abort-input   Aborts a pipeline input
  copy-job      Copy a Jenkins job
  create-job    Create a Jenkins job
  env           Print the current environment
  help          Help about any command
  trigger-input Triggers a pipeline input

Flags:
      --config string   config file (default is $HOME/.jenk.yaml)
  -h, --help            help for jenk
  -t, --toggle          Help message for toggle

Use "jenk [command] --help" for more information about a command.
```
