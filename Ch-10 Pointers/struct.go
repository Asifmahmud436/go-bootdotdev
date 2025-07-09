package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

func getMessageText(a *Analytics,b Message){
	(*a).MessagesTotal++
	if b.Success {
		(*a).MessagesSucceeded++
	} else{
		(*a).MessagesFailed++
	}
}
