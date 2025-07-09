package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

func updateBalance(a *customer,t transaction) error{
	if t.transactionType!=transactionDeposit && t.transactionType!=transactionWithdrawal{
		return errors.New("unknown transaction type")
	}
	if t.transactionType == transactionDeposit{
		(*a).balance += t.amount
		return nil
	}
	if t.transactionType == transactionWithdrawal{
		if((*a).balance<t.amount){
			return errors.New("insufficient funds")
		}else{
			(*a).balance -= t.amount
		}

	}
	return nil
}
