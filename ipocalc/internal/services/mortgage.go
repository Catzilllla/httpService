package services

import (
	"math"

	"github.com/Catzilllla/httpService/ipocalc/internal/models"
)

func CalculateMortgage(prog models.JsonProgram, req models.JsonRequest) (models.JsonAggregate, error) {
	// var newProgramm jsonProgram

	// Определение ставки
	var rate float64
	switch {
	case prog.Military:
		rate = 9
	case prog.Salary:
		rate = 8
	case prog.Base:
		rate = 10
	}

	loanSum := 123.3243
	monthlyPayment := 1000.344
	overpayment := 324.434
	lastPaymentStr := "dsfsd"

	return models.JsonAggregate{
		Rate:            rate,
		LoanSum:         math.Round(loanSum),
		MonthlyPayment:  math.Round(monthlyPayment),
		Overpayment:     math.Round(overpayment),
		LastPaymentDate: lastPaymentStr,
	}, nil
}
