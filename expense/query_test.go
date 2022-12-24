package expense

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
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

		give := MyDB{&Dummy{}}
		want := "MyError"
		_, err := give.QueryExpense_(1)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), want)
	})

}
