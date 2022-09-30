package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	println("Server is Running!")
	e.Logger.Fatal(e.Start(":3333"))
}
