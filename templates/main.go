package main

import (
	"os"
	"text/template"
)

func main() {
	// Defining the data to pass
	type User struct {
		Name string
	}
	user := User{"James"}
	t, err := template.ParseFiles("tmpls/text.txt")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
