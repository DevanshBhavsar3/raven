package repository_test

import (
	"context"
	"testing"

	repository "github.com/DevanshBhavsar3/raven/internal/repository/conversation"
	"github.com/DevanshBhavsar3/raven/internal/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateConversation(t *testing.T) {
	ctx := context.Background()
	db := tests.GetMySQLTestDB(ctx, t)

	r := repository.NewConversationRepository(db)

	tcs := []struct {
		title string
		name  string
	}{
		{
			title: "Create Test Conversation",
			name:  "Test Conversation 1",
		},
	}

	for _, tc := range tcs {
		conversation, err := r.CreateConversation(tc.name, "1")
		require.NoError(t, err)

		assert.Equal(t, conversation.Name, tc.name)
		assert.Equal(t, conversation.UserID, "1")
	}
}

func TestGetConversationByID(t *testing.T) {
	ctx := context.Background()
	db := tests.GetMySQLTestDB(ctx, t)

	_ = repository.NewConversationRepository(db)
}
