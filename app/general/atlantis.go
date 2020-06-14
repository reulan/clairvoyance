package general

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// atlantis.yaml dir parsing
// https://www.runatlantis.io/docs/repo-level-atlantis-yaml.html#reference

// All structs copied directly from Atlantis to ensure that the repo config atlantis.yaml can be properly accessed.
// https://github.com/runatlantis/atlantis/blob/master/server/events/yaml/valid/repo_cfg.go

type AtlantisConfiguration struct {
	// Version is the version of the atlantis YAML file.
	Version       int
	Projects      []Project
	Workflows     map[string]Workflow
	Automerge     bool
	ParallelApply bool
	ParallelPlan  bool
}

type Project struct {
	Dir               string
	Workspace         string
	Name              *string
	WorkflowName      *string
	TerraformVersion  *version.Version
	Autoplan          Autoplan
	ApplyRequirements []string
}


type Autoplan struct {
	WhenModified []string
	Enabled      bool
}

type Stage struct {
	Steps []Step
}

type Step struct {
	StepName  string
	ExtraArgs []string
	// RunCommand is either a custom run step or the command to run
	// during an env step to populate the environment variable dynamically.
	RunCommand string
	// EnvVarName is the name of the
	// environment variable that should be set by this step.
	EnvVarName string
	// EnvVarValue is the value to set EnvVarName to.
	EnvVarValue string
}

type Workflow struct {
	Name  string
	Apply Stage
	Plan  Stage
}
