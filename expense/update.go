package expense

import "github.com/lib/pq"

func UpdateRowById(db DBQuery, ex Expense, id int) (Expense, error) {
	stmt, err := db.Prepare("UPDATE expenses SET title=$2 amount=$3 note=$4 tags=$5  where id=$1")
	// ex := Expense{}
	if err != nil {
		return ex, err
	}

	_, err = stmt.Exec(id, ex.Id, ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return ex, err
	}
	return ex, nil

}

func (db MyDB) UpdateRowById_(ex Expense, id int) (Expense, error) {
	stmt, err := db.Prepare("UPDATE expenses SET title=$2 amount=$3 note=$4 tags=$5  where id=$1")
	// ex := Expense{}
	if err != nil {
		return ex, err
	}

	_, err = stmt.Exec(id, ex.Id, ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return ex, err
	}
	return ex, nil

}
