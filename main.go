package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rollbar/rollbar-go"
)

var rollbarToken = "2f19b776460246abb83cbe4682c74fad"

func main() {

	rollbar.SetToken(rollbarToken)
	rollbar.SetEnvironment("production")

	rollbar.Info("Message body goes here")
	rollbar.Wait()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		rollbar.Info("GET /")
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))

}
