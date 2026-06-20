package repository

import (
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

func (r *ConversationRepository) GetConversationByID(id int64) (*models.Conversation, error) {
	query := `
		SELECT id, name, user_id, created_at, updated_at
		FROM conversations
		WHERE id = :id;
	`

	conversation := &models.Conversation{}

	rows, err := r.db.NamedQuery(query, map[string]any{
		"id": id,
	})
	if err != nil {
		return nil, fmt.Errorf("error getting conversation by id: %v", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.StructScan(conversation); err != nil {
			return nil, fmt.Errorf("error scanning conversation: %v", err)
		}
	}

	return conversation, nil
}

func (r *ConversationRepository) CreateConversation(name string, userID string) (*models.Conversation, error) {
	query := `
		INSERT INTO conversations (name, user_id)
		VALUES (:name, :user_id);
	`
	result, err := r.db.NamedExec(query, map[string]any{
		"name":    name,
		"user_id": userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating conversation: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert id: %v", err)
	}

	return r.GetConversationByID(id)
}
