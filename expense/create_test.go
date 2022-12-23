package expense_test

import (
	"errors"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
)

type Dummy struct{}

func (d Dummy) QueryRow(query string, args ...any) expense.Row {
	return DummyRow{}
}

type DummyRow struct{}

func (r DummyRow) Err() error {
	return errors.New("Dummie error")
}
func (r DummyRow) Scan(dest ...any) error {
	return nil
}

type Stub struct{}

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

	t.Run("TestCreateShouldReturnError", func(t *testing.T) {
		give := expense.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}
		want := "Dummie error"
		testdb := Dummy{}

		_, err := expense.Create(testdb, give)

		if err == nil {
			t.Errorf("want error %s got nill ", want)
		}

	})
}
