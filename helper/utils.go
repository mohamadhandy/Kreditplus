package helper

import (
	"fmt"
	"strings"
)

func HitungJumlahBunga(otr float64, persentaseBunga float64) float64 {
	return otr * (persentaseBunga / 100)
}

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

func CalculateBatasKredit(gaji float64) float64 {
	batasKredit := 3 * gaji
	return batasKredit
}
