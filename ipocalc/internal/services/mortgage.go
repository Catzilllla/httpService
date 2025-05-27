package services

import (
	"ipocalc/ipocalc/internal/models"
	"math"
	"time"
)

func CalculateMortgage(requestData models.JsonRequest) (models.JsonAggregate, error) {
	/* Определение ставки */
	var rate float64
	switch {
	case requestData.Program.Military:
		rate = 9
	case requestData.Program.Salary:
		rate = 8
	case requestData.Program.Base:
		rate = 10
	}

	loanMonth := requestData.Months

	/* Переводим годовую процентную ставку в месячную */
	monthlyRate := rate / 12
	loanSum := requestData.ObjectCost - requestData.InitialPayment

	/* аннуитентный ежемесячный платеж */
	pm := loanSum * (monthlyRate * math.Pow(1+monthlyRate, float64(loanMonth))) / (math.Pow(1+monthlyRate, float64(loanMonth)) - 1)

	/* переплата за весь срок кредита */
	totalPayment := pm * float64(loanMonth)
	overpayment := totalPayment - loanSum

	/* рассчитываем дату последнего платежа */
	startDate := time.Now()
	lastPaymentStr := startDate.AddDate(0, loanMonth, 0)

	return models.JsonAggregate{
		Rate:            rate,
		LoanSum:         math.Round(loanSum),
		MonthlyPayment:  math.Round(pm),
		Overpayment:     math.Round(overpayment),
		LastPaymentDate: string(lastPaymentStr.Format("02-01-2006 15:04:05")),
	}, nil
}
