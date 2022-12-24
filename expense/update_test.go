package expense_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
)

type Dummy struct {
}
type MyError struct {
	err string
}

func (m MyError) Error() string {
	return m.err
}
func (d Dummy) QueryRow(query string, args ...any) *sql.Row {
	return &sql.Row{}
}
func (d Dummy) Prepare(query string) (*sql.Stmt, error) {
	return nil, MyError{"MyError"}
}
func (d Dummy) Close() error {
	return MyError{}
}
func TestUpdateOneRow(t *testing.T) {
	t.Run(" PrepareReturnErrorShouldReturnError", func(t *testing.T) {

		give := &Dummy{}
		want := "MyError"
		ex := expense.Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}

		fmt.Printf("expense =>%v\n", ex)
		fmt.Printf("give =>%v\n", give)
		_, err := expense.QueryExpense(give, 1)

		if err.Error() != want {
			t.Errorf("expect: %v got: %v", want, err.Error())
		}
	})

}
