package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/aRaimaiRu/assessment/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type Err struct {
	Message string `json:"message"`
}

func main() {

	db := &handler.MyHandler{&expense.MyDB{
		expense.InitDB(),
	}}

	defer db.Close()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.POST("/expenses", db.HandlerCreate)
	e.GET("/expenses/:id", db.GetExpenseHandle)
	e.PUT("/expenses/:id", db.UpdateExpenseHandler)
	e.GET("/expenses", db.GetAllExpenses)
	go func() {
		if err := e.Start(os.Getenv("PORT")); err != nil && err != http.ErrServerClosed { // Start server
			e.Logger.Fatal("shutting down the server")
		}
	}()

	gracefulShutdown(e)

}
func gracefulShutdown(e *echo.Echo) {
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
