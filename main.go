package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var customers = make(map[int]*Customer)

func ShowTitle() {
	fmt.Println("Amartha Billing Engine CLI")
}

func ShowHelp() {
	fmt.Println("Help - Available Commands")
	fmt.Println("  help")
	fmt.Println("  add_customer <name>")
	fmt.Println("  show_customer <customer_id>")
	fmt.Println("  make_loan <customer_id>")
	fmt.Println("  make_payment <customerId> <loan_id> <amount>")
	fmt.Println("  get_outstanding <customerId> <loan_id>")
	fmt.Println("  get_missed_weeks_of_payment <customerId> <loan_id>")
	fmt.Println("  is_delinquent <customerId> <loan_id>")
	fmt.Println("  show_report")
	fmt.Println("  exit")
}

func RequiredParamSufficient(tokens []string, sumOfParam int) bool {
	return len(tokens) == sumOfParam+1
}

func AddCustomer(tokens []string) {
	if !RequiredParamSufficient(tokens, 1) {
		fmt.Println("Usage: add_customer <name>")
		return
	}

	name := tokens[1]
	id := len(customers) + 1
	customer := NewCustomer(id, name)
	customers[id] = customer
	fmt.Println("AddCustomer: success -> created")
}

func ShowCustomer(tokens []string) {
	if !RequiredParamSufficient(tokens, 1) {
		fmt.Println("Usage: show_customer <customer_id>")
		return
	}

	customerId, _ := strconv.Atoi(tokens[1])
	customer, exists := customers[customerId]
	if !exists {
		fmt.Println("ShowCustomer: failed -> Customer not found")
		return
	}
	ShowReportHeader()
	ShowReportValue(customer)
}

func MakeLoan(tokens []string) {
	if !RequiredParamSufficient(tokens, 1) {
		fmt.Println("Usage: add_loan <customer_id>")
		return
	}

	customerId, _ := strconv.Atoi(tokens[1])
	customer, exists := customers[customerId]
	if !exists {
		fmt.Println("MakeLoan: failed -> Customer not found")
		return
	}

	customer.MakeLoan()
}

func MakePayment(tokens []string) {
	if !RequiredParamSufficient(tokens, 3) {
		fmt.Println("Usage: make_payment <customerId> <loanId> <amount>")
		return
	}

	customerId, _ := strconv.Atoi(tokens[1])
	customer := customers[customerId]
	if customer == nil {
		fmt.Println("MakePayment: failed -> Customer not found")
		return
	}

	loanId, _ := strconv.Atoi(tokens[2])
	loan := customer.Loans[loanId]
	if loan == nil {
		fmt.Println("MakePayment: failed -> Loan not found")
		return
	}

	amount, _ := strconv.ParseFloat(tokens[3], 64)
	loan.MakePayment(amount)
}

func GetOutstanding(tokens []string) {
	if !RequiredParamSufficient(tokens, 2) {
		fmt.Println("Usage: get_outstanding <customerId> <loanId>")
		return
	}

	customerId, _ := strconv.Atoi(tokens[1])
	customer := customers[customerId]
	if customer == nil {
		fmt.Println("GetOutstanding: failed -> Customer not found")
		return
	}

	loanId, _ := strconv.Atoi(tokens[2])
	loan := customer.Loans[loanId]
	if loan == nil {
		fmt.Println("GetOutstanding: failed -> Loan not found")
		return
	}

	fmt.Printf("| %-3s | %-10s | %-7s | %-15s |\n",
		"id", "name", "loanId", "outstanding")
	fmt.Printf("| %-3d | %-10s | %-7d | %-15.2f |\n",
		customer.ID, customer.Name, loan.ID, loan.GetOutstanding())
}

func GetMissedWeeksOfPayment(tokens []string) {
	if !RequiredParamSufficient(tokens, 2) {
		fmt.Println("Usage: get_missed_weeks_of_payment <customerId> <loanId>")
		return
	}

	customerId, _ := strconv.Atoi(tokens[1])
	customer := customers[customerId]
	if customer == nil {
		fmt.Println("GetMissedWeeksOfPayment: failed -> Customer not found")
		return
	}

	loanId, _ := strconv.Atoi(tokens[2])
	loan := customer.Loans[loanId]
	if loan == nil {
		fmt.Println("GetMissedWeeksOfPayment: failed -> Loan not found")
		return
	}

	fmt.Printf("| %-3s | %-10s | %-7s | %-13s |\n",
		"id", "name", "loanId", "missedWeeks")
	fmt.Printf("| %-3d | %-10s | %-7d | %-13d |\n",
		customer.ID, customer.Name, loan.ID, loan.GetMissedWeeksOfPayment())
}

func IsDelinquent(tokens []string) {
	if !RequiredParamSufficient(tokens, 2) {
		fmt.Println("Usage: is_delinquent <customerId> <loanId>")
		return
	}

	customerId, _ := strconv.Atoi(tokens[1])
	customer := customers[customerId]
	if customer == nil {
		fmt.Println("IsDelinquent: failed -> Customer not found")
		return
	}

	loanId, _ := strconv.Atoi(tokens[2])
	loan := customer.Loans[loanId]
	if loan == nil {
		fmt.Println("IsDelinquent: failed -> Loan not found")
		return
	}

	fmt.Printf("| %-3s | %-10s | %-7s | %-13s |\n",
		"id", "name", "loanId", "delinquent")
	fmt.Printf("| %-3d | %-10s | %-7d | %-13t |\n",
		customer.ID, customer.Name, loan.ID, loan.IsDelinquent())
}

func ShowReport(tokens []string) {
	if !RequiredParamSufficient(tokens, 0) {
		fmt.Println("Usage: show_report")
		return
	}

	ShowReportHeader()
	for _, customer := range customers {
		ShowReportValue(customer)
	}
}

func ShowReportHeader() {
	fmt.Printf("| %-3s | %-10s | %-7s | %-5s | %-10s | %-15s | %-15s |\n",
		"id", "name", "loanId", "week", "status", "amount", "outstanding")
}

func ShowReportValue(customer *Customer) {
	if len(customer.Loans) == 0 {
		fmt.Printf("| %-3d | %-10s | %-7s | %-5s | %-10s | %-15s | %-15s |\n",
			customer.ID, customer.Name, "", "", "", "", "")
	} else {
		for _, loan := range customer.Loans {
			fmt.Printf("| %-3d | %-10s | %-7d | %-5d | %-10s | %-15.2f | %-15.2f |\n",
				customer.ID, customer.Name, loan.ID, loan.Week, loan.Status, loan.Amount, loan.Outstanding)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	ShowTitle()
	ShowHelp()

	for {
		fmt.Print("> ")
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "exit" {
			break
		}
		tokens := strings.Split(line, " ")
		if len(tokens) == 0 {
			continue
		}

		switch tokens[0] {
		case "help":
			ShowHelp()
		case "add_customer":
			AddCustomer(tokens)
		case "show_customer":
			ShowCustomer(tokens)
		case "make_loan":
			MakeLoan(tokens)
		case "make_payment":
			MakePayment(tokens)
		case "get_outstanding":
			GetOutstanding(tokens)
		case "get_missed_weeks_of_payment":
			GetMissedWeeksOfPayment(tokens)
		case "is_delinquent":
			IsDelinquent(tokens)
		case "show_report":
			ShowReport(tokens)
		default:
			fmt.Println("Unknown command, use 'help' to see available commands")
		}
	}
}
