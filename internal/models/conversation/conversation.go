package models

type Conversation struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	UserID    string `db:"user_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
