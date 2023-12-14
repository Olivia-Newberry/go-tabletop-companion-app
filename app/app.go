package app

import (
	"log"

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
	router := echo.New()

	router.Static("/static", "static")

	router.GET("/", handlers.HelloPage)
	router.POST("hello", handlers.HelloName)

	log.Println("handlers are setup")

	if err := database.Migrate(); err != nil {
		log.Fatalf("cannot initialize the database: %v", err)
	}

	log.Println("database is setup")

    router.Logger.Fatal(router.Start(":1323"))
}
