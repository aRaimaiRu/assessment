package expense_test

import (
	"database/sql"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
)

type Dummy struct{}

func (d Dummy) QueryRow(query string, args ...any) *sql.Row {
	return &sql.Row{}
}

//	type error interface {
//		Error() string
//	}
type Myerror struct{}

func (d Myerror) Error() string {
	return "dmummy error"
}

func TestCreateShouldReturnExpense(t *testing.T) {
	// t.Run("TestCreateShouldReturnExpense", func(t *testing.T) {
	// 	give := expense.Expense{
	// 		Title:  "strawberry smoothie",
	// 		Amount: 79,
	// 		Note:   "night market promotion discount 10 bath",
	// 		Tags:   []string{"food", "beverage"},
	// 	}
	// 	want := expense.Expense{
	// 		Title:  "strawberry smoothie",
	// 		Amount: 79,
	// 		Note:   "night market promotion discount 10 bath",
	// 		Tags:   []string{"food", "beverage"},
	// 	}
	// 	testdb := Stub{want}

	// 	got, err := expense.Create(testdb, give)

	// 	if err != nil {
	// 		t.Errorf("Error : %v", err)
	// 	}
	// 	if got.Title != want.Title {
	// 		t.Errorf("want %s got %s", want.Title, got.Title)
	// 	}
	// 	if got.Note != want.Note {
	// 		t.Errorf("want %s got %s", want.Note, got.Note)
	// 	}

	// })

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
