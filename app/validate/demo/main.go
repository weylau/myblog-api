package main

import (
	"fmt"
	"myblog-api/app/validate"
)

func main() {
	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "",
		Tagline:  "This tagline is way too long.",
		Tagline2: "1",
	}
	check,_ := validate.Default()
	if ret := check.CheckStruct(user); !ret{
		fmt.Println(check.GetAllError())
		fmt.Println(check.GetOneError())
	}


}

