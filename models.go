package main

import "text/template"

type user struct {
	UserName  string
	Password  []byte
	FirstName string
	LastName  string
}

var (
	tpl *template.Template 
	dbSessions = make(map[string]string) // session id, user id 
)

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}