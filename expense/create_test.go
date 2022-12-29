package expense

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type ThisError struct {
	S string
}

func (m ThisError) Error() string {
	return m.S
}

func TestCreate(t *testing.T) {
	t.Run("TestCreateRow", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		exp := Expense{
			Title:  "smoothie updated",
			Amount: 99,
			Note:   "night market promotion discount 99 bath",
			Tags:   []string{"beverage"},
		}
		mock_Err := ThisError{S: "mock Create Error"}
		defer db.Close()
		mock.ExpectQuery("INSERT INTO expenses \\(title, amount, note ,tags \\) values \\(\\$1, \\$2, \\$3, \\$4\\) RETURNING id, title, amount, note, tags;").
			WithArgs(exp.Title, exp.Amount, exp.Note, pq.Array(&exp.Tags)).WillReturnError(mock_Err)

		mydb := &MyDB{db}
		_, err = mydb.Create(exp)
		assert.Equal(t, mock_Err, err)

	})

}
