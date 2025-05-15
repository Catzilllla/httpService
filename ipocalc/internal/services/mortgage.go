package services

import (
	"math"
)

func calculateMortgage() (jsonAggregate, error) {
	var newProgramm jsonProgram

	rate := 1
	loanSum := 1
	monthlyPayment := 1
	overpayment := 1
	lastPaymentStr := 1

	return jsonAggregate{
		Rate:            rate,
		LoanSum:         math.Round(loanSum),
		MonthlyPayment:  math.Round(monthlyPayment),
		Overpayment:     math.Round(overpayment),
		LastPaymentDate: lastPaymentStr,
	}, nil
}
