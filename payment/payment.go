package payment

// Payment (Kontrak) / Signature of Method --> Jenis Pembayaran
type Payment interface {
	Pay(amount int64) error
}
