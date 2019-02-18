package payment

import "fmt"

type Gopay struct {
	PhoneNumber int64
}

func (g *Gopay) Pay(amount int64) error {
	fmt.Printf("Phone Number %v pays %v Rupiah using Go Pay\n", g.PhoneNumber, amount)
	return nil
}

func NewGopayAccount(phoneNumber int64) *Gopay {
	return &Gopay{phoneNumber}
}
