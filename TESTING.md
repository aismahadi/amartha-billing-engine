# Amartha Billing Engine

To run the program, please execute this command in terminal :
`go run ./main`

For testing purpose, the programmer made some commands executable from terminal.
Those commands are :

**help**
- show all available commands

**add_customer <name>**
- add customer to Amartha Billing Engine

**show_customer <customer_id>**
- show list of loans of a specified customer

**make_loan <customer_id>**
- makes a loan of a specified customer

**make_payment <customerId> <loan_id> <amount>**
- makes a payment of a specified amount toward the loan

**get_outstanding <customerId> <loan_id>**
- get customer outstanding amount

**get_missed_weeks_of_payment <customerId> <loan_id>**
- get customer misses consecutive weekly payments

**is_delinquent <customerId> <loan_id>**
- check if customer misses 2 consecutive weekly payments or not

**show_report**
- show all customer loans data

**exit**
- exit Amartha Billing Engine

### Example :
```
add_customer Ismahadi
make_loan 1
make_payment 1 1 110000
make_payment 1 1 0
make_payment 1 1 0
get_outstanding 1 1
get_missed_weeks_of_payment 1 1
is_delinquent 1 1
make_payment 1 1 330000
is_delinquent 1 1
add_customer Resyada
make_loan 2
make_payment 2 1 110000
make_payment 2 1 0
make_payment 2 1 0
is_delinquent 2 1
make_payment 2 1 110000
make_payment 2 1 220000
make_payment 2 1 330000
add_customer Pratama
make_loan 3
make_payment 3 1 110000
make_payment 3 1 0
make_payment 3 1 0
make_payment 3 1 0
make_payment 3 1 0
make_payment 3 1 550000
make_loan 3
make_payment 3 2 0
make_payment 3 2 220000
make_payment 3 2 110000
is_delinquent 3 2
make_payment 3 2 0
make_payment 3 2 0
is_delinquent 3 2
show_report
show_customer 1
show_customer 2
show_customer 3
make_payment 3 2 330000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 110000
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
make_payment 3 2 0
show_report
make_payment 3 2 1100000
show_report
```
You also can copy and paste those example commands in terminal after the program started