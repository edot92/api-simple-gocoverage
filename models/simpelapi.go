package models

import "github.com/edot92/simple-gocoverage/lib"

// StructParamDeret param input array
type StructParamDeret struct {
	ParamDeret []int `json:"param_deret"`
}

// temp for generated doc swagger
type tempResultFromLib struct {
	lib.ResultLib `json:"simpelapi"`
}

// ResSuccses standar format, can custom here :)
type ResSuccses struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// GetSimpelAPI param must be array numeric type []int
func GetSimpelAPI(paramArr []int) (lib.ResultLib, error) {
	results, err := lib.CariDeretMax(paramArr)
	if err != nil {
		return results, err
	}
	return results, nil
}
