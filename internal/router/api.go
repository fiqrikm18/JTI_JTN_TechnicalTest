package router

import (
	"JTI_JTN_TechnicalTest/internal/controller"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	phoneNumberController := controller.NewPhoneNumberController()
	e.POST("/create", phoneNumberController.InsertPhoneNumber)
}
