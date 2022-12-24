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
func (db MyDB) QueryExpense_(id int) (Expense, error) {
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

func (db MyDB) QueryAllExpenses() ([]Expense, error) {
	stmt, err := db.Prepare("SELECT * FROM expenses")
	ex_arr := []Expense{}
	if err != nil {
		return ex_arr, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return ex_arr, err
	}

	for rows.Next() {
		ex := Expense{}

		err = rows.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
		if err != nil {
			return ex_arr, err
		}
		ex_arr = append(ex_arr, ex)
	}

	return ex_arr, nil

}
