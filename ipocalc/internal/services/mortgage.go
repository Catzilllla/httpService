package services

import (
	"ipocalc/ipocalc/internal/models"
	"math"
	"time"
)

func CalculateMortgage(requestData models.JsonRequest, myID int) (models.JsResult, error) {
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

	return models.JsResult{
		ID: myID,
		Params: models.JsonRequest{
			ObjectCost:     requestData.ObjectCost,
			InitialPayment: requestData.InitialPayment,
			Months:         requestData.Months,
			Program: models.JsonProgram{
				Salary:   requestData.Program.Salary,
				Military: requestData.Program.Military,
				Base:     requestData.Program.Base,
			},
		},
		Aggregates: models.JsonAggregate{
			Rate:            rate,
			LoanSum:         loanSum,
			MonthlyPayment:  pm,
			Overpayment:     overpayment,
			LastPaymentDate: string(lastPaymentStr.Format("02-01-2006")),
			// LastPaymentDate: string(lastPaymentStr.Format("02-01-2006 15:04:05")),
		},
	}, nil
}

// {"id":0,"params":{"object_cost":5000000,"initial_payment":1000000,"months":240,
// "program":{"salary":true,"military":false,"base":false}},
// "aggregates":{"rate":8,"loan_sum":4000000,"monthly_payment":2666666.6666666665,"overpayment":636000000,"last_payment_date":"29-05-2045"}}, got
// {"id":0,"params":{"object_cost":5000000,"initial_payment":1000000,"months":240,
// "program":{"salary":true,"military":false,"base":false}},
// "aggregates":{"rate":8,"loan_sum":4000000,"monthly_payment":2666666.6666666665,"overpayment":636000000,"last_payment_date":"29-05-2045"}}
