package handlers

import (
	"context"

	errors "github.com/Olivia-Newberry/go-tabletop-companion-app/templates/errors"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, tpl templ.Component) error{
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return tpl.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func renderServerError(ctx echo.Context, stackTrace string) error {
	component := errors.ServerError("500 Internal Error" , stackTrace)
	return component.Render(context.Background(), ctx.Response().Writer)
}

func renderValidationError(ctx echo.Context, message string) error {
	component := errors.ValidationError("400 Bad Request" , message)
	return component.Render(context.Background(), ctx.Response().Writer)
}

func renderNotFoundError(ctx echo.Context, message string) error {
	component := errors.ValidationError("404 Not found" , message)
	return component.Render(context.Background(), ctx.Response().Writer)
}
