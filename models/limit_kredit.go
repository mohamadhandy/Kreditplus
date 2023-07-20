package models

type LimitKredit struct {
	ID          int     `gorm:"column:id_limit" json:"id_limit"`
	IDKonsumen  int     `gorm:"column:id_konsumen" json:"id_konsumen"`
	Tenor       int     `gorm:"column:tenor" json:"tenor"`
	BatasKredit float64 `gorm:"column:batas_kredit" json:"batas_kredit"`
}

func (LimitKredit) TableName() string {
	return "LimitKredit"
}
