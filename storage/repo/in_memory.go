package repo

// Sotrage In memory
type InMemoryStorageI interface {
	Set(key, value string) error
	SetWithTTl(key, value string, seconds int) error
	Get(key string) (interface{}, error)
}
