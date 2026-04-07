package socialnetwork

// UseCase interface for social network operations
type UseCase interface {
	Execute(command string)
	Query(query string) string
}
