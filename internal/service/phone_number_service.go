package service

import (
	"JTI_JTN_TechnicalTest/internal/config"
	"JTI_JTN_TechnicalTest/internal/model"
	"JTI_JTN_TechnicalTest/internal/repository"
	"reflect"
	"strconv"
)

type PhoneNumberService struct {
	repo *repository.PhoneNumberRepository
}

func NewPhoneNumberService(conn *config.DatabaseConn) *PhoneNumberService {
	repo := repository.NewPhoneNumberRepository(conn.DB)
	return &PhoneNumberService{repo}
}

func (service *PhoneNumberService) Create(request interface{}) error {
	if reflect.TypeOf(request).Kind() == reflect.Slice {
		req := request.([]model.PhoneNumberRequest)
		var phoneNumbers []model.PhoneNumber

		for _, request := range req {
			lastChar, _ := strconv.Atoi(string(request.PhoneNumber[len(request.PhoneNumber)-1]))
			numberType := 0 // 0 ganjil, 1 genap
			if lastChar%2 == 0 {
				numberType = 1
			}
			phoneNumber := model.PhoneNumber{PhoneNumber: request.PhoneNumber, Provider: request.Provider, Type: numberType}
			phoneNumbers = append(phoneNumbers, phoneNumber)
		}

		err := service.repo.Insert(phoneNumbers)
		if err != nil {
			return err
		}
	} else {
		req := request.(model.PhoneNumberRequest)
		lastChar, _ := strconv.Atoi(string(req.PhoneNumber[len(req.PhoneNumber)-1]))
		numberType := 0 // 0 ganjil, 1 genap
		if lastChar%2 == 0 {
			numberType = 1
		}
		phoneNumber := model.PhoneNumber{PhoneNumber: req.PhoneNumber, Provider: req.Provider, Type: numberType}
		err := service.repo.Insert(&phoneNumber)
		if err != nil {
			return err
		}
	}

	return nil
}
