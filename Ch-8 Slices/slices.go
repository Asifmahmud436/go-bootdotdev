package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	firstResult := [3]string{primary,secondary,tertiary}
	secondResult := [3]int{len(primary),len(primary)+len(secondary),len(primary)+len(secondary)+len(tertiary)}
	return firstResult,secondResult 
}
