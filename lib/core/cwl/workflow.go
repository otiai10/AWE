package cwl

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mitchellh/mapstructure"
	//"os"
	//"reflect"
	//"strings"
)

type Workflow struct {
	Inputs       []InputParameter          `yaml:"inputs"`
	Outputs      []WorkflowOutputParameter `yaml:"outputs"`
	Id           string                    `yaml:"id"`
	Steps        []WorkflowStep            `yaml:"steps"`
	Requirements []Requirement             `yaml:"requirements"`
	Hints        []Requirement             `yaml:"hints"` // TODO Hints may contain non-requirement objects. Give warning in those cases.
	Label        string                    `yaml:"label"`
	Doc          string                    `yaml:"doc"`
	CwlVersion   CWLVersion                `yaml:"cwlVersion"`
	Metadata     map[string]interface{}    `yaml:"metadata"`
}

func (w *Workflow) GetClass() string { return "Workflow" }
func (w *Workflow) GetId() string    { return w.Id }
func (w *Workflow) SetId(id string)  { w.Id = id }
func (w *Workflow) is_CWL_minimal()  {}
func (w *Workflow) is_Any()          {}
func (w *Workflow) is_process()      {}

func GetMapElement(m map[interface{}]interface{}, key string) (value interface{}, err error) {

	for k, v := range m {
		k_str, ok := k.(string)
		if ok {
			if k_str == key {
				value = v
				return
			}
		}
	}
	err = fmt.Errorf("Element \"%s\" not found in map", key)
	return
}

func NewWorkflow(object CWL_object_generic, collection *CWL_collection) (workflow Workflow, err error) {

	// convert input map into input array

	inputs, ok := object["inputs"]
	if ok {
		err, object["inputs"] = NewInputParameterArray(inputs)
		if err != nil {
			return
		}
	}

	outputs, ok := object["outputs"]
	if ok {
		object["outputs"], err = NewWorkflowOutputParameterArray(outputs)
		if err != nil {
			return
		}
	}

	// convert steps to array if it is a map
	steps, ok := object["steps"]
	if ok {
		err, object["steps"] = CreateWorkflowStepsArray(steps, collection)
		if err != nil {
			return
		}
	}

	requirements, ok := object["requirements"]
	if ok {
		object["requirements"], err = CreateRequirementArray(requirements)
		if err != nil {
			return
		}
	}

	//switch object["requirements"].(type) {
	//case map[interface{}]interface{}:
	// Convert map of outputs into array of outputs
	//	object["requirements"], err = CreateRequirementArray(object["requirements"])
	//	if err != nil {
	//		return
	//	}
	//case []interface{}:
	//	req_array := []Requirement{}

	//	for _, requirement_if := range object["requirements"].([]interface{}) {
	//		switch requirement_if.(type) {

	//		case map[interface{}]interface{}:

	//			requirement_map_if := requirement_if.(map[interface{}]interface{})
	//			requirement_data_if, xerr := GetMapElement(requirement_map_if, "class")

	//			if xerr != nil {
	///				err = fmt.Errorf("Not sure how to parse Requirements, class not found")
	//				return
	//			}

	//			switch requirement_data_if.(type) {
	//			case string:
	//				requirement_name := requirement_data_if.(string)
	//				requirement, xerr := NewRequirement(requirement_name, requirement_data_if)
	//				if xerr != nil {
	//					err = fmt.Errorf("error creating Requirement %s: %s", requirement_name, xerr.Error())
	//					return
	//				}
	//				req_array = append(req_array, requirement)
	//			default:
	//				err = fmt.Errorf("Not sure how to parse Requirements, not a string")
	//				return

	//			}
	//		default:
	//			err = fmt.Errorf("Not sure how to parse Requirements, map expected")
	//			return

	//		} // end switch

	//	} // end for
	//
	//object["requirements"] = req_array
	//}
	fmt.Printf("......WORKFLOW raw")
	spew.Dump(object)
	//fmt.Printf("-- Steps found ------------") // WorkflowStep
	//for _, step := range elem["steps"].([]interface{}) {

	//	spew.Dump(step)

	//}

	err = mapstructure.Decode(object, &workflow)
	if err != nil {
		err = fmt.Errorf("error parsing workflow class: %s", err.Error())
		return
	}
	fmt.Printf(".....WORKFLOW")
	spew.Dump(workflow)
	return
}
