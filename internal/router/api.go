package router

import (
	"JTI_JTN_TechnicalTest/internal/controller/api"
	"github.com/labstack/echo/v4"
)

func RegisterRouter(e *echo.Echo) {
	phoneNumberController := api.NewPhoneNumberController()
	e.POST("/create", phoneNumberController.InsertPhoneNumber)
	e.DELETE("/delete/:id", phoneNumberController.DeletePhoneNumber)
	e.PUT("/update/:id", phoneNumberController.UpdatePhoneNumber)
	e.GET("/get", phoneNumberController.GetAllPhoneNumber)
}
