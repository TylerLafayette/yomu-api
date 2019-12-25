package cache

// Driver represents a driver from which a cache can be accessed.
type Driver interface {
	// Get gets a single item by Key.
	Get(key Key) (Value, error)
	// Set assigns the provided value to the key (updates or creates).
	Set(key Key, value Value) error
	// Delete deletes an item by key (invalidation of a cache item).
	Delete(key Key) error
}
