package expense_test

import (
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
)

func TestCreateShouldReturnExpense(t *testing.T) {
	t.Run("TestCreateShouldReturnExpense", func(t *testing.T) {
		give := expense.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		want := expense.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		testdb := expense.InitDB()

		got := expense.Create(testdb, give)

		if got.Title != want.Title {
			t.Errorf("want %s got %s", want.Title, got.Title)
		}
		if got.Note != want.Note {
			t.Errorf("want %s got %s", want.Note, got.Note)
		}

	})
}
