package accounts

import (
	"bank/clients"
	"fmt"
)

type CheckingAccount struct {
	Owner                       clients.Client // lower-case first letters make an attribute private. Upper-case first letters make an attribute public
	AgencyNumber, AccountNumber int
	balance                     float64
}

// the (c *CheckingAccount) makes it so this becomes a method for the CheckingAccount struct. It's just like an extension function in Kotlin.
func (c *CheckingAccount) Withdraw(withdrawal float64) string {
	canWithdraw := withdrawal > 0 && withdrawal <= c.balance
	if canWithdraw {
		c.balance -= withdrawal
		return fmt.Sprint("Successful withdrawal of ", withdrawal, " dollars by "+c.Owner.Name, ". Your new balance is ", c.balance, " dollars.")
	} else {
		return "Invalid value or insufficient funds"
	}
}

func (c *CheckingAccount) Deposit(depositValue float64) (string, float64) {

	if depositValue > 0 {
		c.balance += depositValue
		return "Deposit was successful", c.balance
	}
	return "Deposit value is equal or less than 0", c.balance

}

func (c *CheckingAccount) Transfer(transferValue float64, targetAccount *CheckingAccount) { // when taking structs as parameters, we need to point to them

	canTransfer := transferValue > 0 && transferValue <= c.balance

	if canTransfer {
		c.Withdraw(transferValue)
		targetAccount.Deposit(transferValue)
	}
}

func (c *CheckingAccount) GetBalance() float64 {
	return c.balance
}
