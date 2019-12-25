package cache

// Service provides a service with abstracted access to a simple
// key-value cache.
type Service struct {
	driver Driver
}

// NewService creates and returns a new *Service.
func NewService(driver Driver) *Service {
	return &Service{driver}
}

// Get retrieves an item from the cache.
func (s *Service) Get(key string) (interface{}, error) {
	return s.driver.Get(Key(key))
}

// Set sets an item in the cache. If an item with the key is already
// present, it will be overwrtitten.
func (s *Service) Set(key string, value interface{}) error {
	return s.driver.Set(Key(key), Value(value))
}

// Delete deletes a specific item in the cache by its key.
func (s *Service) Delete(key string) error {
	return s.driver.Delete(Key(key))
}
