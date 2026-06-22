package conversation

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ConversationRepository struct {
	db *sqlx.DB
}

func NewConversationRepository(db *sqlx.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) CreateConversation(ctx context.Context, userID string, name string) (int64, error) {
	query := `
		INSERT INTO conversations (name, user_id)
		VALUES (:name, :user_id);
	`
	result, err := r.db.NamedExecContext(ctx, query, map[string]any{
		"name":    name,
		"user_id": userID,
	})
	if err != nil {
		return 0, fmt.Errorf("error creating conversation: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error getting last insert id: %v", err)
	}

	return id, nil
}

func (r *ConversationRepository) GetAllConversations(ctx context.Context, userID string) ([]Conversation, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at
		FROM conversations
		WHERE user_id = ?;
	`

	conversations := []Conversation{}

	err := r.db.SelectContext(ctx, &conversations, query, userID)
	if err != nil {
		return nil, fmt.Errorf("error getting all conversations: %v", err)
	}

	return conversations, nil
}

func (r *ConversationRepository) GetConversationByID(ctx context.Context, userID string, id int64) (*Conversation, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at
		FROM conversations
		WHERE id = ? AND user_id = ?;
	`

	conversation := &Conversation{}
	err := r.db.GetContext(ctx, conversation, query, id, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		return nil, fmt.Errorf("error getting conversation by id: %v", err)
	}

	return conversation, nil
}

func (r *ConversationRepository) DeleteConversation(ctx context.Context, userID string, id int64) error {
	query := `
		DELETE FROM conversations
		WHERE id = ? AND user_id = ?;
	`

	result, err := r.db.ExecContext(ctx, query, id, userID)
	if err != nil {
		return fmt.Errorf("error deleting conversation: %v", err)
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return fmt.Errorf("no conversation found with id %d for user %s", id, userID)
	}

	return nil
}
