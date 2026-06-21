package repository

import (
	"context"
	"fmt"

	models "github.com/DevanshBhavsar3/raven/internal/models/conversation"
	"github.com/jmoiron/sqlx"
)

type ConversationRepository struct {
	db *sqlx.DB
}

func NewConversationRepository(db *sqlx.DB) *ConversationRepository {
	return &ConversationRepository{db: db}
}

func (r *ConversationRepository) GetConversationByID(ctx context.Context, id int64) (*models.Conversation, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at
		FROM conversations
		WHERE id = :id;
	`

	conversation := &models.Conversation{}

	err := r.db.SelectContext(ctx, conversation, query, map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting conversation by id: %v", err)
	}

	return conversation, nil
}

func (r *ConversationRepository) CreateConversation(ctx context.Context, c *models.Conversation) (*models.Conversation, error) {
	query := `
		INSERT INTO conversations (name, user_id, created_at)
		VALUES (:name, :user_id, :created_at);
	`
	result, err := r.db.NamedExecContext(ctx, query, map[string]any{
		"name":       c.Name,
		"user_id":    c.UserID,
		"created_at": c.CreatedAt,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating conversation: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %v", err)
	}
	c.ID = id

	return c, nil
}
