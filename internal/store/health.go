package store

import (
	"errors"
	"go.etcd.io/bbolt"
)

func (s *Store) Ready() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("store closed")
	}
	return s.db.View(func(*bbolt.Tx) error { return nil })
}
