package cppcomponents

import (
	"fmt"
	"strings"
)

// Function ...
type Function struct {
	Name        string
	ReturnType  string
	Parameters  []Parameter
	ConstString string
}

// NewFunction .. Constructor
func NewFunction(pureVirtualFunctionLine string) *Function {
	f := Function{}

	// Remove "virtual " from string
	pureVirtualFunctionLine = strings.TrimPrefix(pureVirtualFunctionLine, "virtual ")

	// Parse function Name and return type
	returnTypeAndName := strings.Split(pureVirtualFunctionLine, "(")[0]
	returnTypeAndNameSlice := strings.Split(returnTypeAndName, " ")

	f.Name = returnTypeAndNameSlice[len(returnTypeAndNameSlice)-1]
	f.ReturnType = strings.Join(returnTypeAndNameSlice[:len(returnTypeAndNameSlice)-1], " ")

	// Parse parameter list
	rawParameters := strings.Split(strings.Split(pureVirtualFunctionLine, ")")[0], "(")[1]
	if len(rawParameters) > 0 {
		//TODO: This comma won't work for templated arguments such as QMap<QString, QString>
		rawParametersSlice := strings.Split(rawParameters, ",")
		for _, rawParameterString := range rawParametersSlice {
			f.Parameters = append(f.Parameters, *NewParameter(rawParameterString))
		}
	}

	// Parse function const-ness
	f.ConstString = ""
	if strings.Contains(strings.Split(pureVirtualFunctionLine, ")")[1], "const") {
		f.ConstString = " const"
	}

	return &f
}

/// TODO: Make \t resource configurable (3 spaces, 4 spaces?)
func (f Function) Declaration() string {
	return fmt.Sprintf("\t%v %v(%v)%v override;", f.ReturnType, f.Name, f.allParameters(), f.ConstString)
}

func (f Function) Definition(classScope string) string {
	return fmt.Sprintf("%v %v::%v(%v)%v\n{\n}", f.ReturnType, classScope, f.Name, f.allParameters(), f.ConstString)
}

func (f Function) allParameters() string {
	parametersString := ""
	separator := ""
	for i, p := range f.Parameters {
		if i > 0 {
			separator = ", "
		}
		parametersString += fmt.Sprintf("%v%v", separator, p.toString())
	}
	return parametersString
}

// func splitParameters(rawParameterString string) string {
// 	var commaIndices []int
// 	// Template depth increases for each "<" found, and decreases for each ">"
// 	templateDepth := 0
// 	for pos, char := range rawParameterString {

// 	}

// 	return "butt"

// }
