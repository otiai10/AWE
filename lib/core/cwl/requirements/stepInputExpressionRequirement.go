package requirements

import (
	"github.com/mitchellh/mapstructure"
)

// Indicate that the workflow platform must support the valueFrom field of WorkflowStepInput.

//http://www.commonwl.org/v1.0/Workflow.html#StepInputExpressionRequirement
type StepInputExpressionRequirement struct {
	//Class         string `yaml:"class"`
}

func (c StepInputExpressionRequirement) GetClass() string { return "StepInputExpressionRequirement" }
func (c StepInputExpressionRequirement) GetId() string    { return "None" }

func NewStepInputExpressionRequirement(original interface{}) (r *StepInputExpressionRequirement, err error) {
	var requirement StepInputExpressionRequirement
	r = &requirement
	err = mapstructure.Decode(original, &requirement)
	return
}
