package expense_test

import (
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
)

type Stub struct {
}

func (d Stub) QueryRow(query string, args ...any) expense.Row {
	return StunRow{}
}

type StunRow struct {
}

func (r StunRow) Err() error {
	return nil
}
func (r StunRow) Scan(dest ...any) error {
	return nil
}

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
		testdb := Stub{}

		got, err := expense.Create(testdb, give)

		if err != nil {
			t.Errorf("Error : %v", err)
		}
		if got.Title != want.Title {
			t.Errorf("want %s got %s", want.Title, got.Title)
		}
		if got.Note != want.Note {
			t.Errorf("want %s got %s", want.Note, got.Note)
		}

	})
}
