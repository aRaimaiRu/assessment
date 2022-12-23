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
	return errors.New("Dummy error")
}
func (r DummyRow) Scan(dest ...any) error {
	return errors.New("Dummy error")
}

type Stub struct {
	expense.Expense
}

func (d Stub) QueryRow(query string, args ...any) expense.Row {

	return StubRow{d.Expense}
}

type StubRow struct {
	expense.Expense
}

func (r StubRow) Err() error {
	return nil
}
func (r StubRow) Scan(dest ...any) error {

	value, ok := dest[0].(*int)
	if !ok {
		return errors.New("Error ID")
	} else {
		*value = r.Id
	}

	value_string, ok := dest[1].(*string)
	if !ok {
		return errors.New("Error Title")
	} else {
		*value_string = r.Title
	}

	value_float, ok := dest[2].(*float32)
	if !ok {
		return errors.New("Error Amount")
	} else {
		*value_float = r.Amount
	}

	value_string, ok = dest[3].(*string)
	if !ok {
		return errors.New("Error Note")
	} else {
		*value_string = r.Note
	}

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
		testdb := Stub{want}

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
		want := "Dummy error"
		testdb := Dummy{}

		_, err := expense.Create(testdb, give)

		if err == nil {
			t.Errorf("want error %s got nill ", want)
		}

	})
}
