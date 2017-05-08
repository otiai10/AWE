package cwl

// CWLType - CWL basic types: int, string, boolean, .. etc
// http://www.commonwl.org/v1.0/CommandLineTool.html#CWLType
// null, boolean, int, long, float, double, string, File, Directory
type CWLType interface {
	CWL_minimal_interface
	is_CWLType()
	//is_CWL_minimal()
}

type CWLType_Impl struct{}

func (c *CWLType_Impl) is_CWL_minimal()               {}
func (c *CWLType_Impl) is_CWLType()                   {}
func (c *CWLType_Impl) is_CommandInputParameterType() {}

func NewCWLType(native interface{}) (cwl_type CWLType, err error) {

	switch native.(type) {
	case int:
		native_int := native.(int)

		cwl_type = Int{Value: native_int}
	case string:
		native_str := native.(string)

		cwl_type = String{Value: native_str}
	case bool:
		native_bool := native.(bool)

		cwl_type = Boolean{Value: native_bool}

	default:
		err = fmt.Errorf("(NewAny) Type unknown")

	}
	return

}
