package expense

import (
	"github.com/lib/pq"
)

func (db MyDB) Create(expense Expense) (Expense, error) {
	rows := db.QueryRow("INSERT INTO expenses (title, amount, note ,tags ) values ($1, $2, $3, $4) RETURNING id, title, amount, note, tags;",
		expense.Title, expense.Amount, expense.Note, pq.Array(&expense.Tags))
	ex := Expense{}
	err := rows.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return Expense{}, err
	}
	return ex, nil
}
