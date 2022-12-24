//go:build integration
// +build integration

package expense_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/stretchr/testify/assert"
)

func TestGetExpense(t *testing.T) {
	e := seedExpense(t)

	var lastest expense.Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(e.Id)), nil)
	err := res.Decode(&lastest)

	assert.Nil(t, err)
	assert.Equal(t, lastest.Title, e.Title)
	assert.Equal(t, lastest.Amount, e.Amount)
	assert.Equal(t, lastest.Note, e.Note)
}
