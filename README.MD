# Amartha Take-Home Test – Billing Engine

This project is a solution to the Billing Engine problem from the Amartha backend engineering take-home test.

I chose this problem because it focuses on core backend logic and system design.


## Problem Summary

The idea is to simulate a loan system where:

- A borrower gets a 50-week loan of Rp 5,000,000 with a flat interest of 10% per year.
- Weekly payment is Rp 110,000 (so total is Rp 5,500,000).
- Borrower must pay the exact weekly amount or skip it.
- The system should be able to track outstanding balance and check if the borrower is delinquent.
- A borrower is considered delinquent if they miss 2 or more consecutive payments.


## What I Built

This repo contains:

- A `Loan` struct that holds loan info, payment status, and provides methods like:
  - `MakePayment(week int)`
  - `GetOutstanding()`
  - `IsDelinquent()`

- A simulated in-memory DB (`LoanRepo`) to manage multiple loans
- A simulated message queue (`EventBus`) to emit async events when a payment is made
- A main program (`main.go`) that simulates 10 payment attempts (some valid, some invalid)


## Assumptions

- Only exact payments are allowed (no partial or early/late overpayments)
- Payments can only be made by week number (1 to 50)
- No persistence or real MQ used — everything is simulated in-memory
- If a week is skipped, it's assumed the borrower will pay later in exact order


## Concurrency

- I used `sync.Mutex` to protect access to the `Loan` object
- `main.go` simulates concurrent payments using goroutines and `sync.WaitGroup`
- `EventBus.EmitAsync` is also done with goroutines to simulate async message publishing


## How to Run

```bash
go run main.go
