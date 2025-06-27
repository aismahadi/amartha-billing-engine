# Amartha Billing Engine

Billing Engine (System Design and Abstraction)
We are building a Billing Engine to support our Loan Engine.

**Main Features**
- Generate Loan Repayment Schedule:
  - Provide a detailed weekly repayment schedule for each loan
  - Specifying how much needs to be paid and when
- Track Outstanding Amount:
  - Maintain and update the remaining unpaid loan balance (outstanding amount)
- Identify Delinquent Borrowers:
  - Determine whether a borrower has missed payments and should be classified as delinquent

**Loan Details**
- Each loan has a duration of 50 weeks
- The loan principal is Rp 5,000,000
- A flat interest rate of 10% per annum is applied
- Therefore, the total repayment amount is Rp 5,500,000 (Rp 5,000,000 + 10% interest)
- The borrower makes fixed weekly payments of: Rp 110,000 per week for 50 weeks

**Notes**
- Borrowers can only pay the exact weekly installment scheduled for that week (no partial or excess payments are allowed), or choose not to pay at all
- We must track the outstanding loan balance, which starts at Rp 5,500,000 and decreases with each payment; by the end of the loan term, the outstanding amount should be Rp 0
- If a borrower misses 2 consecutive weekly payments, they are considered a delinquent borrower
- To make up for missed payments, borrowers must pay the total amount for all missed weeks; for example, if 2 weeks were missed, the borrower must repay 2 full weekly installments

**Required Methods to Implement**
- We require at least the following methods to be implemented:
- GetOutstanding:
  - Returns the current outstanding (unpaid) balance of a loan
  - Returns 0 if the loan is fully paid off
- IsDelinquent:
  - Returns whether the borrower has missed more than 2 consecutive weekly payments
- MakePayment:
  - Makes a payment of a specified amount toward the loan