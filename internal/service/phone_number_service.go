package service

import (
	"JTI_JTN_TechnicalTest/internal/config"
	"JTI_JTN_TechnicalTest/internal/model"
	"JTI_JTN_TechnicalTest/internal/repository"
)

type PhoneNumberService struct {
	repo *repository.PhoneNumberRepository
}

func NewPhoneNumberService(conn *config.DatabaseConn) *PhoneNumberService {
	repo := repository.NewPhoneNumberRepository(conn.DB)
	return &PhoneNumberService{repo}
}

func (service *PhoneNumberService) Create(request model.PhoneNumberRequest) error {
	phoneNumber := model.PhoneNumber{PhoneNumber: request.PhoneNumber, Provider: request.Provider}
	err := service.repo.Insert(phoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (service *PhoneNumberService) CreateBatch(request []model.PhoneNumberRequest) error {
	phoneNumbers := []model.PhoneNumber{}

	for idx, _ := range phoneNumbers {
		request := phoneNumbers[idx]
		phoneNumber := model.PhoneNumber{PhoneNumber: request.PhoneNumber, Provider: request.Provider}
		phoneNumbers = append(phoneNumbers, phoneNumber)
	}

	err := service.repo.Insert(phoneNumbers)
	if err != nil {
		return err
	}

	return nil
}
