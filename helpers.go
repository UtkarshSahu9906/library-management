package main

import "time"


func calculateDueDate(borrowDate time.Time) time.Time {
	return borrowDate.AddDate(0, 0, 14)
}
func calculateFine(dueDate string, returnDate string) float64 {
	due, _ := time.Parse("2006-01-02", dueDate)
	returned, _ := time.Parse("2006-01-02", returnDate)

	diff := returned.Sub(due)
	days := int(diff.Hours() / 24)

	finePerDay := 5.0
	return float64(days) * finePerDay
}