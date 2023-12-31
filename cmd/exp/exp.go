package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  template.HTML // XSS?
}

func main() {
	t, err := template.ParseFiles("home.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{
		Name: "Andrzej",
		Bio:  "<script>alert(!)'</script>",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
