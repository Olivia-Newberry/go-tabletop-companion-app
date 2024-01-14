package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/Olivia-Newberry/go-tabletop-companion-app/database"
	"github.com/Olivia-Newberry/go-tabletop-companion-app/handlers"
	"github.com/a-h/templ"

	"github.com/labstack/echo/v4"
)

func HTML(c echo.Context, comp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return comp.Render(c.Request().Context(), c.Response().Writer)
}

func Start() {
	err := godotenv.Load()
	if err != nil {
	log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	router := echo.New()

	router.Static("/static", "static")

	router.GET("/", handlers.HelloPage)
	router.POST("hello", handlers.HelloName)

	log.Println("handlers are setup")

	if err := database.Migrate(); err != nil {
		log.Fatalf("cannot initialize the database: %v", err)
	}

	log.Println("database is setup")

	router.Logger.Fatal(router.Start(":" + port))
}
