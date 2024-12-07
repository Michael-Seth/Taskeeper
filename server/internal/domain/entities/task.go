package entities

type Task struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Completed   bool   `db:"completed"`
	UserID      int    `db:"user_id"`
}
