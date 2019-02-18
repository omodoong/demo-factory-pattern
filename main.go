package main

import (
	payment "haidlir/learn/factory-pattern/payment"
)

func main() {
	// Akun Gojek Mufti
	var muftiPhoneNumber int64 = 6285715338630
	muftiGopay := payment.NewGopayAccount(muftiPhoneNumber)

	// Akun Debit Mufti
	var muftiDebitNumber int64 = 33756928317236132
	muftiDebit := payment.NewDebitAccount(muftiDebitNumber)

	// Mufti Bayar Mie Ayam 11500
	var hargaMieAyam int64 = 11500
	_ = muftiGopay.Pay(hargaMieAyam)
	_ = muftiDebit.Pay(hargaMieAyam)
}
