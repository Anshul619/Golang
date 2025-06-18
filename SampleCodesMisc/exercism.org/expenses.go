package main

import (
	"errors"
	"log"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {

	out := []Record{}

	for _, r := range in {
		if predicate(r) {
			out = append(out, r)
		}
	}

	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return (r.Day >= p.From && p.To >= r.Day)
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	amount := 0.0

	for _, r := range in {
		if ByDaysPeriod(p)(r) {
			amount += r.Amount
		}
	}

	return amount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	amount := 0.0
	isCat := false

	for _, r := range in {
		if ByDaysPeriod(p)(r) && ByCategory(c)(r) {
			amount += r.Amount
		}

		if ByCategory(c)(r) && !isCat {
			isCat = true
		}
	}

	if !isCat {
		return 0, errors.New("unknown category entertainment")
	}

	return amount, nil
}

func main() {
	records := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
	}

	log.Println(Filter(records, func(r Record) bool { return r.Day == 1 }))

	records1 := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	period := DaysPeriod{From: 1, To: 15}

	log.Println(Filter(records1, ByDaysPeriod(period)))

	records2 := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	log.Println(Filter(records2, ByCategory("groceries")))

	records3 := []Record{
		{Day: 15, Amount: 16, Category: "entertainment"},
		{Day: 32, Amount: 20, Category: "groceries"},
		{Day: 40, Amount: 30, Category: "entertainment"},
	}

	p1 := DaysPeriod{From: 1, To: 30}
	p2 := DaysPeriod{From: 31, To: 60}

	log.Println(TotalByPeriod(records3, p1))
	// => 16

	log.Println(TotalByPeriod(records3, p2))
	// => 50

	records4 := []Record{
		{Day: 1, Amount: 15, Category: "groceries"},
		{Day: 11, Amount: 300, Category: "utility-bills"},
		{Day: 12, Amount: 28, Category: "groceries"},
		{Day: 26, Amount: 300, Category: "university"},
		{Day: 28, Amount: 1300, Category: "rent"},
	}

	log.Println(CategoryExpenses(records4, p1, "entertainment"))
	// => 0, error(unknown category entertainment)

	log.Println(CategoryExpenses(records4, p1, "rent"))
	// => 1300, nil

	log.Println(CategoryExpenses(records4, p2, "rent"))
	// => 0, nil
}
