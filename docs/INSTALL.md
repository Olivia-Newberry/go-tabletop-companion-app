# Installation and Setup

Created with Go, HTMX, echo, templ & tailwindcss

## Pre-requisites

1. Go 1.20+
2. [templ CLI](https://templ.guide/quick-start/installation)
3. [air](https://github.com/cosmtrek/air)
4. C compiler gcc / MinGW (Windows)

## Installing Tools

1. Install templ (requires Go 1.20 or higher)

   ```sh
    go install github.com/a-h/templ/cmd/templ@latest
   ```

2. Install air for hot reloading server on file changes

   ```sh
    go install github.com/cosmtrek/air@latest
   ```

3. SQLx

   ```sh
    go get github.com/jmoiron/sqlx@latest
   ```

4. MinGW
    Windows

    ```sh
    choco install mingw
    ```

    If you don't have choco, install it from [here](https://chocolatey.org/install)

    Linux

    You could install build-essential package

    ```sh
    sudo apt-get install build-essential
    ```

## Setup

1. Set environment variables

   ```sh
    go env -w CGO_ENABLED=1
   ```

## Running

```sh
air
```

Go to `http://127.0.0.1:1312`

## Built with

- [Go](https://go.dev/)
- [echo](https://echo.labstack.com/)
- [HTMX](https://htmx.org/)
- [TEMPL](https://templ.guide/)
- [tailwindcss](https://tailwindcss.com/)
