package middlewares

type Middlewares struct {
	Auth *AuthMiddleware
}

func NewMiddlewares() *Middlewares {
	return &Middlewares{
		Auth: NewAuthMiddleware(),
	}
}
