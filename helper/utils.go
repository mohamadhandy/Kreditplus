package helper

import (
	"fmt"
	"strings"
)

// dan persentase bunga yang diberikan.
func HitungJumlahBunga(otr float64, persentaseBunga float64) float64 {
	return otr * (persentaseBunga / 100)
}

// Fungsi untuk menghitung admin fee dalam transaksi berdasarkan nilai admin fee yang diberikan.
func HitungAdminFee(otr float64) float64 {
	return otr * 0.05
}

func GenerateNomorKontrak(jenisTransaksi string, idKonsumen int, jumlahCicilan int) string {
	result := ""
	if strings.ToLower(jenisTransaksi) == "pembelian" {
		result = "PB"
	} else if strings.ToLower(jenisTransaksi) == "pinjaman" {
		result = "PN"
	}
	result += fmt.Sprintf("%v%v", idKonsumen, jumlahCicilan)
	return result
}
