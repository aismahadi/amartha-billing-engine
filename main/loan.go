package main

import "fmt"

type Loan struct {
	ID          int
	CustomerID  int
	Week        int
	Status      string
	Amount      float64
	Outstanding float64
	Payments    map[int]*Payment
}

var base_loan float64 = 5_000_000
var interest float64 = 10
var total_loan float64 = base_loan + (base_loan * interest / 100)
var weeklyPayment float64 = 110_000

func NewLoan(customer *Customer) *Loan {
	return &Loan{
		ID:          len(customer.Loans) + 1,
		CustomerID:  customer.ID,
		Week:        1,
		Status:      "Promising",
		Amount:      total_loan,
		Outstanding: total_loan,
		Payments:    make(map[int]*Payment),
	}
}

func (loan *Loan) GetOutstanding() float64 {
	var paid float64 = 0
	for _, payment := range loan.Payments {
		paid += payment.Amount
	}
	return loan.Amount - paid
}

func (loan *Loan) GetMissedWeeksOfPayment() int {
	var totalPayments int = len(loan.Payments)
	if totalPayments == 0 {
		return 0
	}
	missed := 0
	for i := totalPayments - 1; i >= 0; i-- {
		if loan.Payments[i].Amount == 0 {
			missed++
		} else {
			break
		}
	}
	return missed
}

func (loan *Loan) IsDelinquent() bool {
	return loan.GetMissedWeeksOfPayment() >= 2
}

func (loan *Loan) MakePayment(amount float64) {
	if loan.Status == "Settled" {
		fmt.Println("MakePayment: failed -> Settled")
		return
	}
	var requiredAmount float64 = float64(loan.GetMissedWeeksOfPayment()+1) * weeklyPayment
	deadline := loan.Week == int(total_loan/weeklyPayment)
	if amount == 0 && deadline {
		fmt.Printf("MakePayment: failed -> This week is the deadline, must be settled (Required %.2f)\n", requiredAmount)
		return
	}
	if amount == 0 || amount == requiredAmount {
		loan.Payments[loan.Week-1] = NewPayment(loan, amount)
		if !deadline {
			loan.Week++
		}
		loan.Outstanding = loan.GetOutstanding()
		if loan.Outstanding == 0.0 {
			loan.Status = "Settled"
		} else if loan.IsDelinquent() {
			loan.Status = "Delinquent"
		} else {
			loan.Status = "Promising"
		}
		if amount == 0 {
			fmt.Println("MakePayment: success -> skipped")
		} else {
			fmt.Println("MakePayment: success -> paid")
		}
		return
	}
	fmt.Printf("MakePayment: failed -> Unmatch Amount (Required %2.f or %.2f, Given %.2f)\n", 0.00, requiredAmount, amount)
}
