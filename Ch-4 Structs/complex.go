package main

type messageToSend struct {
	message    string
	sender     user
	reciepient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.reciepient.name!= "" && mToSend.reciepient.number!= 0 && mToSend.sender.name!= "" && mToSend.sender.number!= 0{
		return true
	}
	return false
}