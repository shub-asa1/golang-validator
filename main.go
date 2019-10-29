package main

import (
	//"fmt"
	validate "validator"
)

type temp struct {
	UserEmail string `validate:"email"`
}

func main() {
	v := temp{
		UserEmail : "shubhgmial.com",
	}
	validate.Struct(v);
//	fmt.Println(validate.Var("shubhgmail.com", "email"))
}
