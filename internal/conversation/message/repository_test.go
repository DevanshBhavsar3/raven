package message_test

import (
	"context"
	"os"
	"testing"

	"github.com/DevanshBhavsar3/raven/internal/conversation"
	"github.com/DevanshBhavsar3/raven/internal/conversation/message"
	"github.com/DevanshBhavsar3/raven/internal/tests"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var db *sqlx.DB

func TestMain(m *testing.M) {
	ctx := context.Background()
	testDB := tests.NewMySQLTestDB(ctx)
	db = testDB.GetDB()

	code := m.Run()

	testDB.Terminate(ctx)

	os.Exit(code)
}

func TestCreateMessage(t *testing.T) {
	tcs := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "create new message without conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := message.NewMessageRepository(db)

				messageID, err := r.CreateMessage(ctx, "Test Message", message.MessageRoleUser, 1)
				require.Error(t, err)

				assert.Equal(t, messageID, int64(0))
			},
		},
		{
			title: "create new message with conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				conversationRepo := conversation.NewConversationRepository(db)
				messageRepo := message.NewMessageRepository(db)

				conversationID, err := conversationRepo.CreateConversation(ctx, "message-test-user-id", "Test Conversation")
				require.NoError(t, err)

				messageID, err := messageRepo.CreateMessage(ctx, "Test Message", message.MessageRoleUser, conversationID)
				require.NoError(t, err)

				assert.NotEqual(t, messageID, int64(0))
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, tc.test)
	}
}

func TestGetMessageByConversationID(t *testing.T) {
	tcs := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "get messages by conversation id",
			test: func(t *testing.T) {
				ctx := t.Context()
				conversationRepo := conversation.NewConversationRepository(db)
				messageRepo := message.NewMessageRepository(db)

				conversationID, err := conversationRepo.CreateConversation(ctx, "message-test-user-id", "Test Conversation 1")
				require.NoError(t, err)

				_, err = messageRepo.CreateMessage(ctx, "Test Message 1", message.MessageRoleUser, conversationID)
				require.NoError(t, err)

				_, err = messageRepo.CreateMessage(ctx, "Test Message 2", message.MessageRoleUser, conversationID)
				require.NoError(t, err)

				messages, err := messageRepo.GetMessagesByConversationID(ctx, conversationID)
				require.NoError(t, err)

				assert.Len(t, messages, 2)
			},
		},
		{
			title: "get messages by conversation id with no messages",
			test: func(t *testing.T) {
				ctx := t.Context()
				conversationRepo := conversation.NewConversationRepository(db)
				messageRepo := message.NewMessageRepository(db)

				conversationID, err := conversationRepo.CreateConversation(ctx, "message-test-user-id", "Test Conversation 2")
				require.NoError(t, err)

				messages, err := messageRepo.GetMessagesByConversationID(ctx, conversationID)
				require.NoError(t, err)

				assert.Len(t, messages, 0)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, tc.test)
	}
}
