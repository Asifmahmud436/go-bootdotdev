package main

func getLast[T any](s []T) T {
	var lastVal T = s[len(s)-1]
	return lastVal
}
