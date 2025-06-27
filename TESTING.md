# Amartha Billing Engine

To run the program, please go to "./main" directory and execute this command in terminal :
`go run *`

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