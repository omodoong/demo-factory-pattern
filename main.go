package main

import (
	payment "haidlir/learn/factory-pattern/payment"
	"log"
)

func main() {
	// Dompet Mufti
	dompetMufti := []payment.Payment{}
	// 1 Gopay
	// 2 Debit
	// 3 Tcash
	dompetMuftiConfig := [][]int64{
		[]int64{1, 6285715338630},
		[]int64{2, 33756928317236132},
		[]int64{3, 6285715338630},
		[]int64{1, 6285715339999},
		[]int64{4, 6285715339999},
	}

	for _, config := range dompetMuftiConfig {
		paymentMethod, err := payment.NewPayment(config[0], config[1])
		if err != nil {
			log.Printf("Payment Method %v not found: %v", config[0], err)
			continue
		}
		dompetMufti = append(dompetMufti, paymentMethod)
	}

	// Mufti Bayar Mie Ayam 11500
	var hargaMieAyam int64 = 11500
	for _, bill := range dompetMufti {
		_ = bill.Pay(hargaMieAyam)
	}
}
