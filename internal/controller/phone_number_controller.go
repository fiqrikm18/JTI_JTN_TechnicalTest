package controller

import (
	"JTI_JTN_TechnicalTest/internal/config"
	"JTI_JTN_TechnicalTest/internal/model"
	"JTI_JTN_TechnicalTest/internal/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"reflect"
	"time"
)

type PhoneNumberController struct {
	service *service.PhoneNumberService
}

func NewPhoneNumberController() *PhoneNumberController {
	dbConn, err := config.NewDatabaseConn()
	if err != nil {
		panic(err)
	}

	phoneService := service.NewPhoneNumberService(dbConn)
	return &PhoneNumberController{
		service: phoneService,
	}
}

func (p *PhoneNumberController) InsertPhoneNumber(c echo.Context) error {
	phoneNumberRequest := model.PhoneNumberRequest{}
	var phoneNumberResponses []model.PhoneNumberRequest
	if err := c.Bind(&phoneNumberRequest); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if reflect.ValueOf(phoneNumberRequest).IsZero() {
		var phoneNumbersRequest []model.PhoneNumberRequest
		for i := 0; i < 25; i++ {
			phoneNumber := generatePhoneNumber()
			phoneProvider := model.PhoneProviders[rand.Intn(3)]

			phoneNumbersRequest = append(phoneNumbersRequest, model.PhoneNumberRequest{
				phoneNumber, phoneProvider,
			})
		}

		err := p.service.Create(phoneNumbersRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		phoneNumberResponses = phoneNumbersRequest
	} else {
		err := p.service.Create(phoneNumberRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		phoneNumberResponses = append(phoneNumberResponses, phoneNumberRequest)
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "phone number created",
		"data":    phoneNumberResponses,
	})
}

func generatePhoneNumber() string {
	rand.Seed(time.Now().UnixNano())
	phoneNumber := "08"

	for i := 1; i < 11; i++ {
		phoneNumber += fmt.Sprintf("%d", rand.Intn(10))
	}

	return phoneNumber
}
