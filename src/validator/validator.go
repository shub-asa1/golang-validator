package validator

import(
	"reflect"
	"runtime/debug"
	"fmt"
)

type trav struct{
	fieldname string
	val interface{}
}

//Var dfdsf
func Var(val interface{},tags string) (newerr error){
	
	traverseVal := &trav{
		fieldname : "",
		val: val,
	} 
	
	result := traverse(traverseVal,tags);
	
	errorr := &err{
		message : "validation failed",
		errStack : string(debug.Stack()),
		fielder : result,
	};
	
	
	newerr = errorr;
	return;
}

//Struct dfd
func Struct(val interface{}){

	fi := reflect.ValueOf(val);
	
	errorr := &serr{
		message : "validation failed",
		errStack : string(debug.Stack()),
	};

	for i := 0;i < fi.NumField() ; i++{
		fieldname :=fi.Type().Field(i).Name;
		tags:= fi.Type().Field(i).Tag.Get("validate");
		val := fi.Field(i).Interface()
		traverseVal := &trav{
			fieldname : fieldname,
			val: val,
		} 
		result := traverse(traverseVal,tags);
		errorr.fielderr = append(errorr.fielderr,result);
		fmt.Println(errorr)
	}


}