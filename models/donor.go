package models

type Donor struct {
	ID                 uint
	Name               string
	UserName           string
	ContactInfo        string
	FixedMonthlyAmount uint
	AssignedAdminId    uint
	AssignedAdmin      *uint `gorm:"foreignKey:AssignedAdminId"`
}
