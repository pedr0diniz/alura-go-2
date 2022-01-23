package accounts

import (
	"bank/clients"
	"fmt"
)

type SavingsAccount struct {
	Owner                                  clients.Client
	AgencyNumber, AccountNumber, Operation int // multiple parameters can be declared in the same line if they're the same data type
	balance                                float64
}

func (c *SavingsAccount) Withdraw(withdrawal float64) string {
	canWithdraw := withdrawal > 0 && withdrawal <= c.balance
	if canWithdraw {
		c.balance -= withdrawal
		return fmt.Sprint("Successful withdrawal of ", withdrawal, " dollars by "+c.Owner.Name, ". Your new balance is ", c.balance, " dollars.")
	} else {
		return "Invalid value or insufficient funds"
	}
}

func (c *SavingsAccount) Deposit(depositValue float64) (string, float64) {

	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit was successful", c.balance
	}
	return "Deposit value is equal or less than 0", c.balance

}

func (c *SavingsAccount) GetBalance() float64 {
	return c.balance
}
