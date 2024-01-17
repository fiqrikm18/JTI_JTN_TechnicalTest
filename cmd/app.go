package cmd

import (
	"JTI_JTN_TechnicalTest/internal/config"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func Run() {
	svr := echo.New()
	svr.Use(middleware.Recover())
	svr.Use(middleware.Logger())
	svr.Use(middleware.Gzip())

	config.RegisterRouter(svr)

	port := os.Getenv("APP_PORT")
	svr.Logger.Fatal(svr.Start(":" + port))
}
