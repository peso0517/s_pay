package main

import (
       "net/http"
       "github.com/labstack/echo"
       "log"
)

func main () {
     e := echo.New()
     e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, World!")
     })
     log.Println("start")
     e.Logger.Fatal(e.Start(":8082"))
}
