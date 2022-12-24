package expense

import (
	"database/sql"
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
func (d Dummy) Close() error {
	return MyError{}
}
func TestQueryOneRow(t *testing.T) {
	t.Run(" PrepareReturnErrorShouldReturnError", func(t *testing.T) {

		give := &Dummy{}
		want := "MyError"

		_, err := QueryExpense(give, 1)

		if err.Error() != want {
			t.Errorf("expect: %v got: %v", want, err.Error())
		}
	})

}
