package controller

import (
	"JTI_JTN_TechnicalTest/internal/config"
	"JTI_JTN_TechnicalTest/internal/service"
	"github.com/labstack/echo/v4"
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
	return c.JSON(200, "")
}

func (p *PhoneNumberController) InsertBatchPhoneNumber(c echo.Context) error {
	return c.JSON(200, "")
}
