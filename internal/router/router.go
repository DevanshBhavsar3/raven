package router

import (
	"github.com/DevanshBhavsar3/raven/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handlers *handlers.Handlers) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/conversations", func(r chi.Router) {
			r.Post("/", handlers.Conversation.CreateConversation)
			r.Get("/", handlers.Conversation.GetAllConversations)
			r.Get("/{conversationID}", handlers.Conversation.GetConversationByID)
			r.Delete("/{conversationID}", handlers.Conversation.DeleteConversation)
		})
	})

	return r
}
