package main

import (
	"chain/api/routes"
	echo "github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
