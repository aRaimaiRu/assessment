package expense

func Create(db DBQuery, expense Expense) (Expense, error) {
	row := db.QueryRow("INSERT INTO expenses (id , title, amount, note ,tags ) values ($1, $2, %3, %4) ;",
		expense.Id, expense.Title, expense.Amount, expense.Note, expense.Tags)

	if row.Err() != nil {
		return Expense{}, row.Err()
	}
	return expense, nil
}
