package message

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type MessageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) CreateMessage(ctx context.Context, content string, role MessageRole, conversationID int64) (int64, error) {
	query := `
		INSERT INTO messages (content, role, conversation_id)
		VALUES (:content, :role, :conversation_id);
	`
	result, err := r.db.NamedExecContext(ctx, query, map[string]any{
		"content":         content,
		"role":            role,
		"conversation_id": conversationID,
	})
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *MessageRepository) GetMessagesByConversationID(ctx context.Context, conversationID int64) ([]Message, error) {
	query := `
		SELECT id, content, role, conversation_id, created_at, updated_at
		FROM messages
		WHERE conversation_id = ?;
	`

	messages := []Message{}

	err := r.db.SelectContext(ctx, &messages, query, conversationID)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
