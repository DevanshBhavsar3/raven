package conversation_test

import (
	"context"
	"os"
	"testing"

	"github.com/DevanshBhavsar3/raven/internal/conversation"
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

func TestCreateConversation(t *testing.T) {
	tcs := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "create new conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				conversationID, err := r.CreateConversation(ctx, "create-test-user-id", "Test Conversation")
				require.NoError(t, err)

				assert.Equal(t, conversationID, int64(1))
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, tc.test)
	}
}

func TestGetAllConversations(t *testing.T) {
	tcs := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "get conversations for a user with no conversations",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				conversations, err := r.GetAllConversations(ctx, "get-conversations-test-user-id")
				require.NoError(t, err)

				assert.Equal(t, len(conversations), 0)
			},
		},
		{
			title: "get conversations after creating a conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				_, err := r.CreateConversation(ctx, "get-conversations-test-user-id", "Test Conversation")
				require.NoError(t, err)

				conversations, err := r.GetAllConversations(ctx, "get-conversations-test-user-id")
				require.NoError(t, err)

				assert.Equal(t, len(conversations), 1)

				assert.Equal(t, conversations[0].Name, "Test Conversation")
			},
		},
		{
			title: "get conversations for a different user",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				_, err := r.CreateConversation(ctx, "get-conversations-test-user2-id", "Test Conversation 2")
				require.NoError(t, err)

				conversations, err := r.GetAllConversations(ctx, "get-conversations-test-user-id")
				require.NoError(t, err)

				assert.Equal(t, len(conversations), 1)
				assert.Equal(t, conversations[0].Name, "Test Conversation")

				conversations, err = r.GetAllConversations(ctx, "get-conversations-test-user2-id")
				require.NoError(t, err)

				assert.Equal(t, len(conversations), 1)
				assert.Equal(t, conversations[0].Name, "Test Conversation 2")
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, tc.test)
	}
}

func TestGetConversationByID(t *testing.T) {
	tcs := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "get conversation by id",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				conversationID, err := r.CreateConversation(ctx, "get-conversation-by-id-test-user-id", "Test Conversation")
				require.NoError(t, err)

				conversation, err := r.GetConversationByID(ctx, "get-conversation-by-id-test-user-id", conversationID)
				require.NoError(t, err)

				assert.Equal(t, conversation.ID, conversationID)
				assert.Equal(t, conversation.Name, "Test Conversation")
			},
		},
		{
			title: "get conversation by id for a different user",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				conversationID, err := r.CreateConversation(ctx, "get-conversation-by-id-test-user2-id", "Test Conversation 2")
				require.NoError(t, err)

				_, err = r.GetConversationByID(ctx, "get-conversation-by-id-test-user-id", conversationID)
				require.Error(t, err)
			},
		},
		{
			title: "get non-existent conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				_, err := r.GetConversationByID(ctx, "get-conversation-by-id-test-user-id", 999)
				require.Error(t, err)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, tc.test)
	}
}

func TestDeleteConversation(t *testing.T) {
	tcs := []struct {
		title string
		test  func(t *testing.T)
	}{
		{
			title: "delete conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				conversationID, err := r.CreateConversation(ctx, "delete-conversation-test-user-id", "Test Conversation")
				require.NoError(t, err)

				err = r.DeleteConversation(ctx, "delete-conversation-test-user-id", conversationID)
				require.NoError(t, err)

				_, err = r.GetConversationByID(ctx, "delete-conversation-test-user-id", conversationID)
				require.Error(t, err)
			},
		},
		{
			title: "delete conversation for a different user",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				conversationID, err := r.CreateConversation(ctx, "delete-conversation-test-user2-id", "Test Conversation 2")
				require.NoError(t, err)

				err = r.DeleteConversation(ctx, "delete-conversation-test-user-id", conversationID)
				require.Error(t, err)

				conversation, err := r.GetConversationByID(ctx, "delete-conversation-test-user2-id", conversationID)
				require.NoError(t, err)

				assert.Equal(t, conversation.ID, conversationID)
				assert.Equal(t, conversation.Name, "Test Conversation 2")
			},
		},
		{
			title: "delete non-existent conversation",
			test: func(t *testing.T) {
				ctx := t.Context()
				r := conversation.NewConversationRepository(db)

				err := r.DeleteConversation(ctx, "delete-conversation-test-user-id", 999)
				require.Error(t, err)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.title, tc.test)
	}
}
