package service

import (
	"JTI_JTN_TechnicalTest/internal/config"
	"JTI_JTN_TechnicalTest/internal/model"
)

type PhoneNumberService struct {
	DBConn *config.DatabaseConn
}

func NewPhoneNumberService(conn *config.DatabaseConn) *PhoneNumberService {
	return &PhoneNumberService{DBConn: conn}
}

func (service *PhoneNumberService) Create(request model.PhoneNumberRequest) error {
	//phoneNumber := model.PhoneNumber{PhoneNumber: request.PhoneNumber, Provider: request.Provider}
	//err := service.DB(phoneNumber)
	//if err != nil {
	//	return err
	//}
}

func (service *PhoneNumberService) CreateBatch(request []model.PhoneNumberRequest) error {
}
