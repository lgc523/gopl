package main

import (
	"log"
	"os"
	"text/template"
)

type Foo struct {
	Name string
	Age  int
}

func main() {
	text := "Foo: Name:{{.Name}},Age: {{.Age}}"
	tpl, err := template.New("simple").Parse(text)
	if err != nil {
		log.Fatalf("parse error %v", err)
	}
	err = tpl.Execute(os.Stdout, &Foo{Name: "spider", Age: 25})
	if err != nil {
		log.Fatalln(err)
	}

}
