package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed == false {
		return len(e.body) * 5
	}
	return len(e.body) * 3
}

func (e email) format() string {
	info := fmt.Sprintf("'%s'| %v",e.body,e.isSubscribed)
	return info
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
