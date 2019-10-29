package validator

import (
	
)

type fielderr struct {
	fieldname string
	err       string
}



type serr struct{
	message string
	errStack string
	fielderr []fielderr
}

type err struct {
	message  string
	errStack string
	fielder fielderr
}
/*
func (e *err) Error() string {

	result := fmt.Sprintf(" errMsg : %s ,\n errStack : %s \n", e.message, e.errStack)
	fielderror := ""
	fielderror += fmt.Sprintf(" fieldname : %s ,\n err : %s \n", i.fieldname, i.err)
	
	return result + fielderror
}
*/
func (e *err)Error() string{
	return "blah";
}

var errMsg = map[string]string{
	"email": "invalid email address",
}
