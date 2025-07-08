package main

func getMessageCosts(messages []string) []float64 {
	mySlice := make([]float64,len(messages))
	for i:=1 ; i<=len(messages);i++{
		index := float64(i)*0.01*float64(len(messages[i-1]))
		mySlice = append(mySlice, index)
	}
	return mySlice
}
