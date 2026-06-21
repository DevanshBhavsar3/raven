package repository_test

import (
	"context"
	"testing"
	"time"

	models "github.com/DevanshBhavsar3/raven/internal/models/conversation"
	repository "github.com/DevanshBhavsar3/raven/internal/repository/conversation"
	"github.com/DevanshBhavsar3/raven/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateConversation(t *testing.T) {
	ctx := context.Background()
	db := tests.GetMySQLTestDB(ctx, t)

	tcs := []struct {
		title string
		test  func(t *testing.T, ctx context.Context, r *repository.ConversationRepository)
	}{
		{
			title: "create new conversation",
			test: func(t *testing.T, ctx context.Context, r *repository.ConversationRepository) {
				conversation, err := r.CreateConversation(ctx, &models.Conversation{
					Name:      "Test Conversation",
					UserID:    "1",
					CreatedAt: time.Now(),
				})

				require.NoError(t, err)

				assert.Equal(t, conversation.ID, int64(1))
				assert.Equal(t, conversation.Name, "Test Conversation")
				assert.Equal(t, conversation.UserID, "1")
			},
		},
	}

	r := repository.NewConversationRepository(db)

	for _, tc := range tcs {
		tc.test(t, ctx, r)
	}
}

func TestGetConversationByID(t *testing.T) {
	ctx := context.Background()
	db := tests.GetMySQLTestDB(ctx, t)

	_ = repository.NewConversationRepository(db)
}
