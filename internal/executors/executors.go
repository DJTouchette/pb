package executors

import (
	"fmt"
	"pb/internal/parser"
)

type ExecutionContext struct {
	Params  map[string]string
	Results map[string]any
}

func NewExecutionContext(params map[string]string) *ExecutionContext {
	return &ExecutionContext{
		Params:  params,
		Results: make(map[string]any),
	}
}

type StepExecutor interface {
	Execute(step parser.Step) (any, error)
}

var registry = map[string]StepExecutor{}

func RegisterExecutor(stepType string, executor StepExecutor) {
	registry[stepType] = executor
}

func ExecuteStep(step parser.Step, ctx *ExecutionContext) (any, error) {
	executor, ok := registry[step.Type]
	if !ok {
		return fmt.Errorf("unknown step type: %s", step.Type), nil
	}
	err, result := executor.Execute(step)
	if err != nil {
		return err, nil
	}

	if step.SaveResults != "" && result != nil {
		ctx.Results[step.SaveResults] = result
	}

	return result, nil
}

func InitExecutors() {
	FsInit()
}
