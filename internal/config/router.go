package config

import (
	"JTI_JTN_TechnicalTest/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	phoneNumberController := controller.PhoneNumberController{}
	e.POST("/create", phoneNumberController.InsertPhoneNumber)
	e.POST("/create_batch", phoneNumberController.InsertBatchPhoneNumber)
}
