package accounts

type BillPayer interface {
	Withdraw(withdrawal float64) string
}
