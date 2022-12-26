package handler

import (
	"net/http"
	"strconv"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/labstack/echo/v4"
)

type MyHandler struct {
	*expense.MyDB
}
type Err struct {
	Message string `json:"message"`
}

func (db MyHandler) HandlerCreate(c echo.Context) error {
	u := expense.Expense{}
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	u, err = db.Create(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusCreated, u)
}

func (db MyHandler) GetExpenseHandle(c echo.Context) error {
	id_param := c.Param("id")
	id, err := strconv.Atoi(id_param)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	u, err := db.QueryExpense_(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, u)

}

func (db MyHandler) UpdateExpenseHandler(c echo.Context) error {
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

	db.UpdateRowById_(e, id)
	return c.JSON(http.StatusOK, e)

}

func (db MyHandler) GetAllExpenses(c echo.Context) error {
	e, err := db.QueryAllExpenses()
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, e)
}
