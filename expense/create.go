package expense

func Create(db DBQuery, expense Expense) (Expense, error) {
	rows := db.QueryRow("INSERT INTO expenses (id , title, amount, note ,tags ) values ($1, $2, %3, %4) ;",
		expense.Id, expense.Title, expense.Amount, expense.Note, expense.Tags)
	ex := Expense{}
	err := rows.Scan(&ex.Id, &ex.Title, &ex.Amount, &ex.Note, &ex.Tags)
	if err != nil {
		return Expense{}, err
	}
	return ex, nil
}
