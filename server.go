package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

var db *expense.MyDB

type Err struct {
	Message string `json:"message"`
}

func main() {

	db = &expense.MyDB{
		expense.InitDB(),
	}
	defer db.Close()
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "OK")
	})

	e.POST("/expenses", handlerCreate)
	e.GET("/expenses/:id", getExpenseHandle)
	e.PUT("/expenses/:id", UpdateExpenseHandler)
	e.GET("/expenses", getAllExpenses)
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

func handlerCreate(c echo.Context) error {
	u := expense.Expense{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	u, err = expense.Create(db, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, u)
}

func getExpenseHandle(c echo.Context) error {
	id_param := c.Param("id")
	id, err := strconv.Atoi(id_param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	u, err := expense.QueryExpense(db, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, u)

}

func UpdateExpenseHandler(c echo.Context) error {
	id_param := c.Param("id")
	id, err := strconv.Atoi(id_param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	e := expense.Expense{}
	err = c.Bind(&e)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	expense.UpdateRowById(db, e, id)
	return c.JSON(http.StatusOK, e)

}

func getAllExpenses(c echo.Context) error {
	e, err := expense.QueryAllExpenses(db)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, e)
}
