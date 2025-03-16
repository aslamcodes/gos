package donor

type Donor struct {
    ID uint
    Name string
    UserName string
    ContactInfo string
    FixedMonthlyAmount string
    AssignedAdminId uint
    AssignedAdmin *uint `gorm:"foreignKey:AssignedAdminId"`
}
