//go:build integration
// +build integration

package expense_test

import (
	"bytes"
	"net/http"
	"strconv"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/stretchr/testify/assert"
)

func TestUpdateExpense(t *testing.T) {
	t.Run(" UpdateExpenseShouldReturnUpdatedExpense", func(t *testing.T) {
		ex := seedExpense(t)
		got := expense.Expense{}
		body := bytes.NewBufferString(
			`{
			"title": "smoothie updated",
			"amount": 99,
			"note": "night market promotion discount 99 bath", 
			"tags": ["beverage"]
		}`,
		)
		want := expense.Expense{
			Title:  "smoothie updated",
			Amount: 99,
			Note:   "night market promotion discount 99 bath",
			Tags:   []string{"beverage"},
		}

		err := request(http.MethodPut, uri("expenses", strconv.Itoa(ex.Id)), body).Decode(&got)
		assert.Nil(t, err)
		assert.Equal(t, got.Title, want.Title)
		assert.Equal(t, got.Amount, want.Amount)
		assert.Equal(t, got.Tags, want.Tags)

	})

}
