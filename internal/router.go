package internal

import (
	"github.com/DevanshBhavsar3/raven/internal/middlewares"
	"github.com/go-chi/chi/v5"
)

func NewRouter(handlers *Handlers) *chi.Mux {
	r := chi.NewRouter()

	middlewares := middlewares.NewMiddlewares()

	r.Route("/api/v1", func(r chi.Router) {
		// Authenticated routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.Auth.Authenticate)

			r.Route("/conversations", func(r chi.Router) {
				r.Post("/", handlers.Conversation.CreateConversation)
				r.Get("/", handlers.Conversation.GetAllConversations)
				r.Route("/{conversationID}", func(r chi.Router) {
					r.Get("/", handlers.Conversation.GetConversationByID)
					r.Delete("/", handlers.Conversation.DeleteConversation)
					r.Route("/messages", func(r chi.Router) {
						r.Post("/", handlers.Conversation.CreateMessage)
						r.Get("/", handlers.Conversation.GetMessagesByConversationID)
					})
				})
			})
		})
	})

	return r
}
