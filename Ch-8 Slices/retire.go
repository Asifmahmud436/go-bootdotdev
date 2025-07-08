package main

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if(plan == planFree){
		return messages[0:2] , nil
	} else if plan == planPro {
		return messages[0:] , nil
	} else{
		return nil , nil
	}
}
