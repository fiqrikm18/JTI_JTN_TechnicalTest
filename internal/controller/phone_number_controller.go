package controller

import (
	"JTI_JTN_TechnicalTest/internal/config"
	"JTI_JTN_TechnicalTest/internal/model"
	"JTI_JTN_TechnicalTest/internal/service"
	"JTI_JTN_TechnicalTest/internal/util"
	"fmt"
	"github.com/gorilla/websocket"
	"os"

	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
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
			phoneNumber := util.GeneratePhoneNumber()
			phoneProvider := model.PhoneProviders[rand.Intn(3)]

			phoneNumbersRequest = append(phoneNumbersRequest, model.PhoneNumberRequest{
				PhoneNumber: phoneNumber, Provider: phoneProvider,
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
		"message": "phone number deleted",
		"data":    phoneNumberResponses,
	})
}

func (p *PhoneNumberController) DeletePhoneNumber(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	phoneNumber, err := p.service.FindNumberById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = p.service.DeleteNumber(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	dial, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://localhost:%s/ws", os.Getenv("APP_PORT")), nil)
	defer func(dial *websocket.Conn) {
		err := dial.Close()
		if err != nil {

		}
	}(dial)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = dial.WriteMessage(websocket.TextMessage, []byte("data_update"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "phone number deleted",
		"data":    phoneNumber,
	})
}

func (p *PhoneNumberController) UpdatePhoneNumber(c echo.Context) error {
	var request model.PhoneNumberRequest
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	_, err := p.service.FindNumberById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	phoneNumber, err := p.service.UpdateNumber(id, request)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "phone number updated",
		"data":    phoneNumber,
	})
}
