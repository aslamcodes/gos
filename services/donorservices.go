package services

import (
	"fmt"
	"gos/models"
	"gos/repositories"
)

func CreateDonor() (any) {
    donor := &models.Donor{Name: "Alice", UserName: "alice@example.com", ContactInfo: "1234567890",FixedMonthlyAmount: 100, AssignedAdminId: 1 }
    err := repositories.CreateDonor(donor)

    if err != nil {
        fmt.Println("Error:", err)
    }

    return "Created successfully"
}
