package message

import "time"

type MessageRole string

const (
	MessageRoleUser  MessageRole = "user"
	MessageRoleAgent MessageRole = "agent"
)

type Message struct {
	ID             int64       `db:"id"`
	Content        string      `db:"content"`
	Role           MessageRole `db:"role"`
	ConversationID int64       `db:"conversation_id"`
	CreatedAt      time.Time   `db:"created_at"`
	UpdatedAt      time.Time   `db:"updated_at"`
}
