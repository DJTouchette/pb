package executors

import (
	"fmt"
	"os"
	"pb/internal/parser"
)

type PrintExecutor struct{}

func (e *PrintExecutor) Execute(step parser.Step) (any, error) {
	message := step.Params["message"].(string)
	fmt.Println(message)

	return message, nil
}

type CwdExecutor struct{}

func (e *CwdExecutor) Execute(step parser.Step) (any, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return cwd, nil
}

func FsInit() {
	RegisterExecutor("print", &PrintExecutor{})
	RegisterExecutor("get_cwd", &CwdExecutor{})
}
