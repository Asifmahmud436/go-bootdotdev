package main

func addEmailsToQueue(emails []string) chan string {
	ch := make(chan string,len(emails))
	for _,v := range emails{
		ch <- v
	}
	return ch
}
