package expense_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aRaimaiRu/assessment/expense"
	"github.com/lib/pq"
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

		_, err := give.UpdateRowById(ex, 1)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), want)

	})

	t.Run(" ExecReturnErrorShouldReturnError", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		exp := expense.Expense{
			Id:     1,
			Title:  "smoothie updated",
			Amount: 99,
			Note:   "night market promotion discount 99 bath",
			Tags:   []string{"beverage"},
		}
		giveid := 1
		mock_Err := expense.ThisError{S: "QueryError"}
		mock.ExpectPrepare("UPDATE expenses SET title=\\$2 amount=\\$3 note=\\$4 tags=\\$5  where id=\\$1").
			ExpectExec().WithArgs(giveid, exp.Id, exp.Title, exp.Amount, exp.Note, pq.Array(&exp.Tags)).WillReturnError(mock_Err)
		mydb := &expense.MyDB{db}
		_, err = mydb.UpdateRowById(exp, giveid)
		fmt.Printf("Error :%s", err.Error())
		assert.Equal(t, mock_Err, err)

	})

}
