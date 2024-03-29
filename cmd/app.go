package cmd

import (
	"JTI_JTN_TechnicalTest/internal/router"
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
	svr.Use(middleware.CORS())

	router.RegisterRouter(svr)
	router.RegisterWebRouter(svr)
	router.StartWebsocket(svr)

	port := os.Getenv("APP_PORT")
	svr.Logger.Fatal(svr.Start(":" + port))
}
