package models

type Patient struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string
	Age       int
	Address   string
	CreatedBy uint
}
