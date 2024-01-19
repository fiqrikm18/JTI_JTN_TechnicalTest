package repository

import (
	"JTI_JTN_TechnicalTest/internal/model"
	"gorm.io/gorm"
)

type PhoneNumberRepository struct {
	DB *gorm.DB
}

func NewPhoneNumberRepository(db *gorm.DB) *PhoneNumberRepository {
	return &PhoneNumberRepository{
		DB: db,
	}
}

func (repo *PhoneNumberRepository) InsertNumber(phoneNumber interface{}) error {
	return repo.DB.Create(phoneNumber).Error
}

func (repo *PhoneNumberRepository) GetAll() (*[]model.PhoneNumber, error) {
	var phoneNumbers []model.PhoneNumber
	err := repo.DB.Find(&phoneNumbers).Order("id desc").Error
	if err != nil {
		return nil, err
	}
	return &phoneNumbers, nil
}

func (repo *PhoneNumberRepository) FindById(id int) (*model.PhoneNumber, error) {
	var phoneNumber model.PhoneNumber
	err := repo.DB.Find(&phoneNumber, id).Error
	if err != nil {
		return nil, err
	}
	return &phoneNumber, nil
}

func (repo *PhoneNumberRepository) DeleteNumber(id int) error {
	err := repo.DB.Delete(&model.PhoneNumber{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (repo *PhoneNumberRepository) UpdateNumber(id int, data model.PhoneNumber) (*model.PhoneNumber, error) {
	var phoneNumber model.PhoneNumber
	err := repo.DB.First(&phoneNumber).Error
	if err != nil {
		return nil, err
	}

	err = repo.DB.Model(&model.PhoneNumber{}).Where("id=?", id).Updates(&model.PhoneNumber{
		PhoneNumber: data.PhoneNumber,
		Provider:    data.Provider,
		Type:        data.Type,
	}).Scan(&phoneNumber).Error
	if err != nil {
		return nil, err
	}

	return &phoneNumber, nil
}
