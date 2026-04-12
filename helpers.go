package main

import (
	"fmt"
	"time"
)


func calculateDueDate(borrowDate time.Time) time.Time {
	return borrowDate.AddDate(0, 0, 14)
}
// calculateFine returns the fine amount for a borrow record
// only charges fine when book is returned after due date
func calculateFine(dueDate string, returnDate string) float64 {
	due, _ := time.Parse("2006-01-02", dueDate)
	returned, _ := time.Parse("2006-01-02", returnDate)

	diff := returned.Sub(due)
	days := int(diff.Hours() / 24)

	// only charge fine if overdue
	if days <= 0 {
		return 0
	}

	finePerDay := 5.0
	return float64(days) * finePerDay
}
func init() {
	fine1 := calculateFine("2026-04-09", "2026-04-07") // 2 days early
	fine2 := calculateFine("2026-04-09", "2026-04-09") // same day
	fine3 := calculateFine("2026-04-09", "2026-04-16") // 7 days late

	fmt.Println("Early return fine:", fine1)   // should be 0
	fmt.Println("On time fine:", fine2)        // should be 0
	fmt.Println("Late return fine:", fine3)    // should be 35
}