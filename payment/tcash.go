package payment

import "fmt"

type Tcash struct {
	PhoneNumber int64
}

func (t *Tcash) Pay(amount int64) error {
	fmt.Printf("Phone Number %v pays %v Rupiah using Tcash\n", t.PhoneNumber, amount)
	return nil
}

func NewTcashAccount(phoneNumber int64) *Tcash {
	return &Tcash{phoneNumber}
}
