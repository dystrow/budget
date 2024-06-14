package income_test

import (
	"testing"
	"time"

	"github.com/dystrow/budget/api/income"
)

func TestGetPaydayJune_24(t *testing.T) {
	payRanges := [2]income.PayRange{{PayDay: 10, PayPeriodStart: 16, PayPeriodEnd: 0}, //0 is last month
		{PayDay: 25, PayPeriodStart: 1, PayPeriodEnd: 15}}

	workDays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}
	rate := income.HourlyRate{Rate: 62.01, DailyHours: 8, WorkDays: workDays}

	res := income.MonthlyPaydays(2024, time.June, rate, payRanges[:])
	t.Logf("amount: %f, date: %v\n", res[0].Gross, res[0].Date)
	t.Logf("amount: %f, date: %v\n", res[1].Gross, res[1].Date)

	if len(res) != 2 {
		t.Fatalf("result for june24 not 2 but %v\n", len(res))
	}

	if res[0].Gross != 5952.96 {
		t.Fatalf("amount: %f, date: %v\n", res[0].Gross, res[0].Date)
	}
	if res[1].Gross != 4960.80 {
		t.Fatalf("amount: %f, date: %v\n", res[1].Gross, res[1].Date)
	}
}

func TestGetPaydayAug23(t *testing.T) {
	payRanges := [2]income.PayRange{{PayDay: 10, PayPeriodStart: 16, PayPeriodEnd: 0}, //0 is last month
		{PayDay: 25, PayPeriodStart: 1, PayPeriodEnd: 15}}

	workDays := []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}
	rate := income.HourlyRate{Rate: 62.01, DailyHours: 8, WorkDays: workDays}

	res := income.MonthlyPaydays(2023, time.August, rate, payRanges[:])
	t.Logf("amount: %f, date: %v\n", res[0].Gross, res[0].Date)
	t.Logf("amount: %f, date: %v\n", res[1].Gross, res[1].Date)

	if len(res) != 2 {
		t.Fatalf("result for june24 not 2 but %v\n", len(res))
	}

	if res[0].Gross != 5456.88 {
		t.Fatalf("amount: %f, date: %v\n", res[0].Gross, res[0].Date)
	}
	if res[1].Gross != 5456.88 {
		t.Fatalf("amount: %f, date: %v\n", res[1].Gross, res[1].Date)
	}
}
