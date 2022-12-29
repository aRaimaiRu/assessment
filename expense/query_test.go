package expense

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
		_, err := give.QueryExpense(1)

		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), want)
	})
	t.Run("PrepareAndScanError", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock_Err := ThisError{S: "QueryOneError"}
		giveid := 1
		mock.ExpectPrepare("SELECT *").
			ExpectQuery().WithArgs(giveid).WillReturnRows().WillReturnError(mock_Err)
		mydb := &MyDB{db}
		_, err = mydb.QueryExpense(giveid)
		assert.Equal(t, mock_Err, err)

	})

}
func TestQueryAllRow(t *testing.T) {
	t.Run("PrepareAndScanError", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		mock_Err := ThisError{S: "QueryError"}
		mock.ExpectPrepare("SELECT *").
			ExpectQuery().WithArgs().WillReturnRows().WillReturnError(mock_Err)
		mydb := &MyDB{db}
		_, err = mydb.QueryAllExpenses()
		assert.Equal(t, mock_Err, err)

	})
}
