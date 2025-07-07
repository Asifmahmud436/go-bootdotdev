package main

import "fmt"

type Formatter interface {
	Format() string
}

type PlainText struct {
	Message string
}
type Bold struct {
	Message string
}
type Code struct {
	Message string
}

func (p PlainText) Format() string {
	return p.Message
}
func (b Bold) Format() string {
	info := fmt.Sprintf("**%s**",b.Message)
	return info;
}
func (c Code) Format() string {
	info := fmt.Sprintf("`%s`",c.Message)
	return info;
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}
