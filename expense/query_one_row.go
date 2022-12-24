package expense

import "github.com/lib/pq"

func QueryExpense(db DBQuery, id int) (Expense, error) {
	stmt, err := db.Prepare("SELECT * FROM expenses where id=$1")
	ex := Expense{}
	if err != nil {
		return ex, err
	}

	row := stmt.QueryRow(id)
	err = row.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return ex, err
	}
	return ex, nil

}
