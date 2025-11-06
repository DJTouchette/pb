package parser

import (
	"github.com/goccy/go-yaml"
	"path/filepath"
	"io/fs"
	"fmt"
	"os"
)

// Playbook represents the top-level playbook definition
type Playbook struct {
    Name              string              `yaml:"name"`
    Version           string              `yaml:"version,omitempty"`
    Description       string              `yaml:"description,omitempty"`
    Tags              []string            `yaml:"tags,omitempty"`
    Owner             string              `yaml:"owner,omitempty"`
    EstimatedDuration string              `yaml:"estimated_duration,omitempty"`
    Parameters        []Parameter         `yaml:"parameters,omitempty"`
    Preconditions     []Precondition      `yaml:"preconditions,omitempty"`
    Steps             []Step              `yaml:"steps"`
    Rollback          []Step              `yaml:"rollback,omitempty"`
    PostActions       []Step              `yaml:"post_actions,omitempty"`
}

// Parameter defines an input parameter for the playbook
type Parameter struct {
    Name        string   `yaml:"name"`
    Type        string   `yaml:"type"`  // string, int, bool, enum
    Required    bool     `yaml:"required,omitempty"`
    Default     string   `yaml:"default,omitempty"`
    Description string   `yaml:"description,omitempty"`
    Values      []string `yaml:"values,omitempty"`  // For enum type
}

// Precondition defines a check that must pass before execution
type Precondition struct {
    Name    string            `yaml:"name"`
    Type    string            `yaml:"type"`  // command, manual, sql
    Command string            `yaml:"command,omitempty"`
    Prompt  string            `yaml:"prompt,omitempty"`
    When    string            `yaml:"when,omitempty"`  // Conditional expression
    Params  map[string]string `yaml:"params,omitempty"`
}

// Step represents a single executable step
type Step struct {
    Name        string            `yaml:"name"`
    Type        string            `yaml:"type"`  // command, sql, aws, docker, api, parallel, etc.
    Description string            `yaml:"description,omitempty"`
    
    // Command execution fields
    Command    string            `yaml:"command,omitempty"`
    WorkingDir string            `yaml:"working_dir,omitempty"`
    Env        map[string]string `yaml:"env,omitempty"`
    
    // Control flow
    When      string  `yaml:"when,omitempty"`
    OnFailure string  `yaml:"on_failure,omitempty"`
    Timeout   string  `yaml:"timeout,omitempty"`
    
    // Nested steps (for parallel execution)
    Steps []Step `yaml:"steps,omitempty"`
    
    // Generic params for different step types
    Params map[string]interface{} `yaml:"params,omitempty"`
    
    // Sub-playbook execution
    Playbook string                 `yaml:"playbook,omitempty"`
    With     map[string]interface{} `yaml:"with,omitempty"`
    
    // Results
    SaveResults string `yaml:"save_results,omitempty"`
}

func ParsePlaybook(yamlContent []byte) (*Playbook, error) {
    var playbook Playbook
    err := yaml.Unmarshal(yamlContent, &playbook)
    if err != nil {
        return nil, fmt.Errorf("failed to parse playbook: %w", err)
    }
    return &playbook, nil
}

func Discover(pwd string) ([]string, error) {
	var files []string
	filepath.WalkDir(pwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, nil;
}

func GetPlaybook(filePath string) (*Playbook, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return nil, err
    }
    
    var playbook Playbook
    if err := yaml.Unmarshal(data, &playbook); err != nil {
        return nil, err
    }
    
    return &playbook, nil
}

func GetPlaybookBasePath() (string, error) {
		pwd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		fp := pwd + "/playbooks"

		return fp, nil
}
