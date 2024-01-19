package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type PhoneNumberController struct{}

func NewPhoneNumberController() *PhoneNumberController {
	return &PhoneNumberController{}
}

func (p *PhoneNumberController) InputPhoneNumber(c echo.Context) error {
	return c.Render(http.StatusOK, "input.html", nil)
}

func (p *PhoneNumberController) OutputPhoneNumber(c echo.Context) error {
	return c.Render(http.StatusOK, "output.html", nil)
}
