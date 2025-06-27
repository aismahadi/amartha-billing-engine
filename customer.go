package main

import "fmt"

type Customer struct {
	ID    int
	Name  string
	Loans map[int]*Loan
}

func NewCustomer(id int, name string) *Customer {
	return &Customer{
		ID:    id,
		Name:  name,
		Loans: make(map[int]*Loan),
	}
}

func (customer *Customer) MakeLoan() {
	customer.Loans[len(customer.Loans)+1] = NewLoan(customer)
	fmt.Println("MakeLoan: success -> created")
}
