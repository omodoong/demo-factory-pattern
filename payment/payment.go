package payment

import "fmt"

const (
	GopayPayment = 1
	DebitPayment = 2
	TcashPayment = 3
	DokuPayment  = 4
)

// Payment (Kontrak) / Signature of Method --> Jenis Pembayaran
type Payment interface {
	Pay(amount int64) error
}

func NewPayment(paymentMethod, accountID int64) (Payment, error) {
	switch paymentMethod {
	case 1:
		return NewGopayAccount(accountID), nil
	case 2:
		return NewDebitAccount(accountID), nil
	case 3:
		return NewTcashAccount(accountID), nil
	default:
		return nil, fmt.Errorf("Payment Belum Dibuat")
	}
}
