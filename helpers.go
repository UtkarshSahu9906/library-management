package main

import "time"


func calculateDueDate(borrowDate time.Time) time.Time {
	return borrowDate.AddDate(0, 0, 14)
}