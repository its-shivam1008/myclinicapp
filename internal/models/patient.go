package models

type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Age       int
	Prescription string
	Address   string
	CreatedBy uint
}
