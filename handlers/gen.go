package gen

//go:generate mockgen -source=internal/transport/http/handlers.go -destination=internal/mock/user.go UserService
