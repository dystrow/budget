package income

import (
	"fmt"
	"time"
)

// example me
// hourly: 62.01
// working Mon-Friday, 8h each
// payed twice a month: 10th, 25th

// example marina twohands
// hourly, tips. Avg 500 a week
// payed weekly on thursday

// example me before
// yearly: 82000
// payed montly on the 31

// example marina now
// flexible payment

// I'm pretty sure about these structs
type PayRange struct {
	PayDay         int
	PayPeriodStart int
	PayPeriodEnd   int
}

type Payday struct {
	Gross    float64
	Net      float64
	comment  string
	Date     time.Time
	Recieved bool
	from     string
}

type PayRate interface {
	calculatePayment(start time.Time, end time.Time) float64
}

type HourlyRate struct {
	Rate       float64
	DailyHours float64
	WorkDays   []time.Weekday
}

func (rate HourlyRate) calculatePayment(start time.Time, end time.Time) float64 {
	amount := 0.0
	for it := start; it.Before(end.AddDate(0, 0, 1)); it = it.AddDate(0, 0, 1) {
		for _, wd := range rate.WorkDays {
			if wd == it.Weekday() {
				amount += rate.Rate * rate.DailyHours
				break
			}
		}
	}
	return amount
}

func MonthlyPaydays(year int, month time.Month, payRate PayRate, payRanges []PayRange) []Payday {
	// bi monthly first-15, 16-last
	monthStart := time.Date(year, month, 0, 0, 0, 0, 0, time.Local)

	//calculate
	payDays := make([]Payday, len(payRanges))
	for i, payRange := range payRanges {

		payPeriodEnd := monthStart.AddDate(0, 0, payRange.PayPeriodEnd)
		payPeriodStart := payPeriodEnd.AddDate(0, 0, payRange.PayPeriodStart-payPeriodEnd.Day())

		fmt.Printf("year: %v, month: %v starts %v, ends %v, payDay: %v", year, month, payPeriodStart, payPeriodEnd, payRange.PayDay)

		payday := monthStart.AddDate(0, 0, payRange.PayDay)
		gross := payRate.calculatePayment(payPeriodStart, payPeriodEnd)

		p := Payday{Gross: gross, Date: payday, Recieved: false}
		payDays[i] = p
	}

	return payDays
}

type Income interface {
	GetPayday(next int) Payday
}

type HourlyJob struct {
	rate float32
}

type FixedJob struct {
}

type IncomeType int

const (
	Hourly IncomeType = iota + 1
	Salary
)

type PaymentType int

const (
	Daily PaymentType = iota + 1
	Weekly
	BiWeekly
	TwiceMonth
	Monthly
	Yearly
	Flexible
)
