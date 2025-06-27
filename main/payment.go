package main

type Payment struct {
	ID     int
	LoanID int
	Week   int
	Amount float64
}

func NewPayment(loan *Loan, amount float64) *Payment {
	return &Payment{
		ID:     len(loan.Payments) + 1,
		LoanID: loan.ID,
		Week:   len(loan.Payments) + 1,
		Amount: amount,
	}
}
