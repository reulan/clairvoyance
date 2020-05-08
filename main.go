package main

import (
    "fmt"
    tfjson "github.com/hashicorp/terraform-json"
    tfexec "github.com/kmoe/terraform-exec"
)

// Figure out what statefiles need to be parsed
var BasePath string = "/Users/mpmsimo/noobshack/"
var ProjectName string = "valorant"
var StateSubPath string = "/terraform"
var workingDir string = BasePath + ProjectName + StateSubPath

func initializeProjectState() () {
    err := tfexec.Init(workingDir)
    if err != nil {
        panic(err)
    }
    fmt.Println("Project has been initialized.")
}

// For each terraform project, parse the state values
func getProjectState() (*tfjson.State) {
    // terraform show
    state, err := tfexec.Show(workingDir)
    if err != nil {
        panic(err)
    }

    return state
}

// Record the output to be sent to Discord
func sendProjectState(projectState *tfjson.State) {
    // Output retrieved data
    fmt.Println("Displaying project: " + ProjectName)
    fmt.Println("Format Version: " + projectState.FormatVersion)
    fmt.Println("Terraform Version: " + projectState.TerraformVersion)
    fmt.Println("State Values: ")
    fmt.Println(projectState.Values)
}

// TODO: Compare statefile to live resources and output drift
func main() {
    initializeProjectState()
    projectState := getProjectState()
    sendProjectState(projectState)
}
