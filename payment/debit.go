package payment

import "fmt"

type Debit struct {
	DebitNumber int64
}

func (d *Debit) Pay(amount int64) error {
	fmt.Printf("Debit Number %v pays %v Rupiah using DebitCard\n", d.DebitNumber, amount)
	return nil
}

func NewDebitAccount(debitnumber int64) *Debit {
	return &Debit{debitnumber}
}
