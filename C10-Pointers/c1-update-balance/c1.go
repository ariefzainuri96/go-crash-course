package c1updatebalance

/*
[Update Balance]

Textio needs a new way to update users account balance.

[Assignment]

Implement the updateBalance function. It should take a customer pointer and a transaction, and return an error. Depending on the transactionType, it should either add or subtract the amount from the customers balance. If the customer does not have enough money, it should return the error insufficient funds. If the transactionType isn’t a withdrawal or deposit, it should return the error unknown transaction type. Otherwise, it should process the transaction and return nil.
*/

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

func updateBalance(customer *customer, transaction transaction) error {
	if transaction.transactionType != transactionDeposit &&
		transaction.transactionType != transactionWithdrawal {
		return errors.New("unknown transaction type")
	}

	if transaction.amount > customer.balance {
		return errors.New("insufficient funds")
	}

	if customer == nil {
		return errors.New("customer is nil")
	}

	if transaction.transactionType == transactionDeposit {
		(*customer).balance += transaction.amount
	} else {
		(*customer).balance -= transaction.amount
	}

	return nil
}
