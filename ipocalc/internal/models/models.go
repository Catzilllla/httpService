package models

// type ApiError struct {
// 	Message string
// }

// func (e *ApiError) Error() string {
// 	return e.Message
// }

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
}

type JsonRequest struct {
	ObjectCost     float64     `json:"object_cost"`
	InitialPayment float64     `json:"initial_payment"`
	Months         int         `json:"months"`
	Program        JsonProgram `json:"program"`
}

type JsonProgram struct {
	Salary   bool `json:"salary"`
	Military bool `json:"military"`
	Base     bool `json:"base"`
}

type JsonAggregate struct {
	Rate            float64 `json:"rate"`
	LoanSum         float64 `json:"loan_sum"`
	MonthlyPayment  float64 `json:"monthly_payment"`
	Overpayment     float64 `json:"overpayment"`
	LastPaymentDate string  `json:"last_payment_date"`
}

type JsResult struct {
	ID         int           `json:"id"`
	Params     JsonRequest   `json:"params"`
	Aggregates JsonAggregate `json:"aggregates"`
}

// {
// 	"result": {
// 	   "params": {                           // запрашиваемые параметры кредита
// 		  "object_cost": 5000000,
// 		  "initial_payment": 1000000,
// 		  "months": 240
// 	   },
// 	   "program": {                          // программа кредита
// 		  "salary": true
// 	   },
// 	   "aggregates": {                       // блок с агрегатами
// 		  "rate": 8,                         // годовая процентная ставка
// 		  "loan_sum": 4000000,               // сумма кредита
// 		  "monthly_payment": 33458,          // аннуитетный ежемесячный платеж
// 		  "overpayment": 4029920,            // переплата за весь срок кредита
// 		  "last_payment_date": "2044-02-18"  // последняя дата платежа
// 	   }
// 	}
//  }
