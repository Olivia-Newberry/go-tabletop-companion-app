package handlers

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"

	templates "github.com/Olivia-Newberry/go-tabletop-companion-app/templates/hello"
)

func HelloPage(ctx echo.Context) error {
	log.Println("hello handler")

	return templates.HelloPage().Render(context.Background(), ctx.Response().Writer)
}

func HelloName(ctx echo.Context) error {
	name := ctx.FormValue("name")
	log.Println("helloName - entered " + name)

	return templates.HelloName(name).Render(context.Background(), ctx.Response().Writer)
}
