package repositories

import (
	"gos/config"
	"gos/models"
)

// Create a new donor
func CreateDonor(donor *models.Donor) error {
    return config.DB.Create(donor).Error
}

// Get donor by ID
func GetDonorByID(id uint) (*models.Donor, error) {
    var donor models.Donor
    if err := config.DB.First(&donor, id).Error; err != nil {
        return nil, err
    }
    return &donor, nil
}

// Get all donors
func GetAllDonors() ([]models.Donor, error) {
    var donors []models.Donor
    if err := config.DB.Find(&donors).Error; err != nil {
        return nil, err
    }
    return donors, nil
}

// Update donor details
func UpdateDonor(id uint, updatedData *models.Donor) error {
    return config.DB.Model(&models.Donor{}).Where("id = ?", id).Updates(updatedData).Error
}

// Delete a donor by ID
func DeleteDonor(id uint) error {
    return config.DB.Delete(&models.Donor{}, id).Error
}
