package router

import (
	"JTI_JTN_TechnicalTest/internal/controller"
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func RegisterWebRouter(e *echo.Echo) {
	t := &Template{templates: template.Must(template.ParseGlob("public/views/*.html"))}
	e.Renderer = t

	phoneNumberController := controller.NewPhoneNumberController()
	e.GET("/input", phoneNumberController.InputPhoneNumber)
	e.GET("/output", phoneNumberController.OutputPhoneNumber)
}
