package repositories

import (
	repositories "github.com/DevanshBhavsar3/raven/internal/repositories/conversation"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	Conversation *repositories.ConversationRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		Conversation: repositories.NewConversationRepository(db),
	}
}
