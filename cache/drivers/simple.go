package drivers

import (
	"errors"

	"github.com/TylerLafayette/yomu-api/cache"
)

// SimpleDriver provides a simple in-memory cache (primarily for dev/testing).
type SimpleDriver struct {
	store map[cache.Key]cache.Value
}

// NewSimpleDriver creates and returns a *SimpleDriver.
func NewSimpleDriver() *SimpleDriver {
	return &SimpleDriver{
		map[cache.Key]cache.Value{},
	}
}

// Get gets a value from the store by key.
func (d *SimpleDriver) Get(key cache.Key) (cache.Value, error) {
	// Get the value from the map.
	value, ok := d.store[key]
	if !ok {
		// Value not found, return error.
		return nil, errors.New("could not find value")
	}

	// Return the value.
	return cache.Value(value), nil
}

// Set sets a value in the store by key.
func (d *SimpleDriver) Set(key cache.Key, value cache.Value) error {
	d.store[key] = value
	return nil
}

// Delete deletes a store value by key.
func (d *SimpleDriver) Delete(key cache.Key) error {
	delete(d.store, key)
	return nil
}
