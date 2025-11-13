package runner

import (
	"fmt"
	"pb/internal/executors"
	"pb/internal/parser"
	"strings"
)

// Build the parameter values map with defaults
func BuildParamValues(playbook *parser.Playbook, cliParams map[string]string) (map[string]string, error) {
	result := make(map[string]string)

	// First, apply defaults from playbook parameters
	for _, param := range playbook.Parameters {
		if param.Default != "" {
			result[param.Name] = param.Default
		}
	}

	// Then override with CLI-provided values
	for key, value := range cliParams {
		result[key] = value
	}

	// Validate required parameters are present
	for _, param := range playbook.Parameters {
		if param.Required {
			if _, exists := result[param.Name]; !exists {
				return nil, fmt.Errorf("required parameter '%s' not provided", param.Name)
			}
		}
	}

	return result, nil
}

func Execute(playbook *parser.Playbook, paramValues map[string]string) error {
	executors.InitExecutors()
	paramsWithDefault, err := BuildParamValues(playbook, paramValues)
	ctx := executors.NewExecutionContext(paramsWithDefault)

	for _, step := range playbook.Steps {
		// Substitute variables before execution
		if err != nil {
			fmt.Println("Parmas defaults errored: ")
			fmt.Println(err)
			return err
		}

		substitutedStep := SubstituteParams(step, ctx)

		if _, err := executors.ExecuteStep(substitutedStep, ctx); err != nil {
			return err
		}
	}
	return nil
}

func SubstituteParams(step parser.Step, ctx *executors.ExecutionContext) parser.Step {
	substituted := step
	substituted.Params = make(map[string]any)

	for key, value := range step.Params {
		if strVal, ok := value.(string); ok {
			substituted.Params[key] = replaceVariables(strVal, ctx)
		} else {
			substituted.Params[key] = value
		}
	}

	return substituted
}

func replaceVariables(text string, ctx *executors.ExecutionContext) string {
	result := text

	// Replace ${param_name} from parameters
	for key, value := range ctx.Params {
		placeholder := fmt.Sprintf("${%s}", key)
		result = strings.ReplaceAll(result, placeholder, value)
	}

	// Replace ${result_name} from saved results
	for key, value := range ctx.Results {
		placeholder := fmt.Sprintf("${%s}", key)
		result = strings.ReplaceAll(result, placeholder, fmt.Sprintf("%v", value))
	}

	return result
}
