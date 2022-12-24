package expense_test

import (
	"database/sql"
	"testing"

	"github.com/aRaimaiRu/assessment/expense"
	"github.com/stretchr/testify/assert"
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
		give := expense.MyDB{&Dummy{}}
		ex := expense.Expense{}
		want := "MyError"

		_, err := give.UpdateRowById_(ex, 1)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), want)

	})

}
