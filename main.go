package main

import (
	"bank/accounts"
	"bank/clients"
	"fmt"
)

func PayBill(account accounts.BillPayer, value float64) {
	account.Withdraw(value)
}

func main() {
	// Owner := "Pedro"
	// AgencyNumber := 589
	// AccountNumber := 123456
	// Balance := 125.50

	// fmt.Println(Owner, AgencyNumber, AccountNumber, Balance)

	// pedrosAccount := accounts.CheckingAccount{
	// 	Owner:         "Pedro",
	// 	AgencyNumber:  589,
	// 	AccountNumber: 123456,
	// 	Balance:       125.50}

	// pedrosAccount2 := accounts.CheckingAccount{
	// 	Owner:         "Pedro",
	// 	AgencyNumber:  589,
	// 	AccountNumber: 123456,
	// 	Balance:       125.50}

	// amandasAccount := accounts.CheckingAccount{"Amanda", 555, 111333, 200.00}

	// fmt.Println(pedrosAccount == pedrosAccount2)
	// fmt.Println(amandasAccount)

	// var janesAccount *accounts.CheckingAccount
	// janesAccount = new(accounts.CheckingAccount) // the use of the pointer above is necessary so Go knows that we're pointing to janesAccount, not to our struct
	// janesAccount.Owner = "Jane"

	// var janesAccount2 *accounts.CheckingAccount
	// janesAccount2 = new(accounts.CheckingAccount) // the use of the pointer above is necessary so Go knows that we're pointing to janesAccount, not to our struct
	// janesAccount2.Owner = "Jane"

	// // fmt.Println(*janesAccount)
	// fmt.Println(&janesAccount, &janesAccount2) // different pointers
	// fmt.Println(*janesAccount, *janesAccount2) // same content

	// marusAccount := accounts.CheckingAccount{}
	// marusAccount.Owner = "Maru"
	// marusAccount.Balance = 500

	// fmt.Println(marusAccount.Withdraw(300))
	// fmt.Println(marusAccount.Balance)

	// fmt.Println(marusAccount.Withdraw(600))
	// fmt.Println(marusAccount.Balance)

	// fmt.Println(marusAccount.Withdraw(-100))
	// fmt.Println(marusAccount.Balance)

	// status, value := marusAccount.Deposit(10000)
	// fmt.Println(status)
	// fmt.Println(value)

	// glaubersAccount := accounts.CheckingAccount{"Glauber", 123, 549821, 100.0}
	// enerisAccount := accounts.CheckingAccount{"Eneri", 123, 324789, 100.0}

	// glaubersAccount.Transfer(55.0, &enerisAccount) // when passing structs as parameters, we need to pass their values (&enerisAccount),
	// // not the object (enerisAccount) or the pointer (*enerisAccount)
	// fmt.Println(glaubersAccount)
	// fmt.Println(enerisAccount)

	pedrosAccount := accounts.CheckingAccount{
		Owner: clients.Client{
			Name:       "Pedro",
			CPF:        "123.123.123-12",
			Occupation: "Developer",
		},
		AgencyNumber:  123,
		AccountNumber: 123456}
	pedrosAccount.Deposit(1000.0)

	fmt.Println(pedrosAccount)

	amanda := clients.Client{
		Name:       "Amanda",
		CPF:        "123.456.789-00",
		Occupation: "Lawyer",
	}

	amandasAccount := accounts.CheckingAccount{
		Owner:         amanda,
		AgencyNumber:  123,
		AccountNumber: 232893}

	fmt.Println(amandasAccount)

	safeAccount := accounts.CheckingAccount{}

	safeAccount.Deposit(100.0)
	fmt.Println(safeAccount)

	amandaSavingsAccount := accounts.SavingsAccount{
		Owner:         amanda,
		AgencyNumber:  123,
		AccountNumber: 232893,
		Operation:     1}
	amandaSavingsAccount.Deposit(2500.0)

	amandaSavingsAccount.Deposit(1500.0)
	amandaSavingsAccount.Withdraw(1500.0)
	fmt.Println(amandaSavingsAccount)

	fmt.Println(pedrosAccount.GetBalance())
	PayBill(&pedrosAccount, 10)
	fmt.Println(pedrosAccount.GetBalance())

	fmt.Println(amandaSavingsAccount.GetBalance())
	PayBill(&amandaSavingsAccount, 50)
	fmt.Println(amandaSavingsAccount.GetBalance())

	billsSlice := []accounts.BillPayer{&pedrosAccount, &amandaSavingsAccount}

	for j := 0; j < len(billsSlice); j++ {
		fmt.Println(billsSlice[j].Withdraw(50.0))
	}

	// in GoLang, interfaces work kind of the same way as Java, but they're implemented implictly.
	// The methods of the interfaces are their contract. Since the BillPayer interface has the method Withdraw(withdrawal float64) string,
	// any class that can see the BillPayer interface and implements a method called Withdraw that takes a float64 and returns a string
	// WILL automatically implement that interface.

}
