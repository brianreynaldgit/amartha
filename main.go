package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/brianreynaldgit/amartha/loan"
)

func main() {
	repo := loan.NewLoanRepo()
	events := loan.NewEventBus()
	loanID := "loan-1001"

	// Create new loan
	if err := repo.CreateLoan(loanID); err != nil {
		panic(err)
	}
	l, _ := repo.GetLoan(loanID)

	var wg sync.WaitGroup

	paymentWeeks := []int{
		1, 2, 2, 51, 3, 4, 5, 5, 6, -1, // simulate various inputs
	}

	for i, week := range paymentWeeks {
		wg.Add(1)
		go func(i, week int) {
			defer wg.Done()

			err := l.MakePayment(week)
			if err != nil {
				fmt.Printf("[Sim %d] Week %d: %v\n", i+1, week, err)
				return
			}

			events.EmitAsync("PaymentMade", map[string]any{
				"loanID": loanID,
				"week":   week,
				"amount": l.WeeklyPayment,
			})
			fmt.Printf("[Sim %d] Week %d payment successful\n", i+1, week)
		}(i, week)
	}

	wg.Wait()
	time.Sleep(500 * time.Millisecond) // wait for async events

	// Simulate missed payments on week 7 and 8
	fmt.Println("--- FINAL STATUS ---")
	fmt.Println("Outstanding:", l.GetOutstanding())
	fmt.Println("Is Delinquent?", l.IsDelinquent())
}
