package expense

import (
	"database/sql"
	"fmt"
	"testing"
)

type Dummy struct {
}
type MyError struct{}

func (m MyError) Error() string {
	return "MyError"
}
func (d Dummy) QueryRow(query string, args ...any) *sql.Row {
	return &sql.Row{}
}
func (d Dummy) Prepare(query string) (*sql.Stmt, error) {
	return nil, MyError{}
}
func TestQueryOneRow(t *testing.T) {
	t.Run(" PrepareReturnErrorShouldReturnError", func(t *testing.T) {

		give := &Dummy{}
		want := "MyError"
		ex := Expense{
			Title:  "strawberry smoothie",
			Amount: 79,
			Note:   "night market promotion discount 10 bath",
			Tags:   []string{"food", "beverage"},
		}

		fmt.Printf("expense =>%v\n", ex)
		fmt.Printf("give =>%v\n", give)
		_, err := QueryExpense(give, 1)

		if err.Error() != want {
			t.Errorf("expect: %v got: %v", want, err.Error())
		}
	})

}
