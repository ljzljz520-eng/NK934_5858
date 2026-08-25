package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
	"waveboard/internal/domain"
)

var buckets = [][]byte{[]byte("waves"), []byte("tasks"), []byte("zones"), []byte("exceptions")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func put[T any](s *Store, b []byte, id string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return bbolt.ErrDatabaseNotOpen
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(id), data) })
}
func get[T any](s *Store, b []byte, id string, out *T) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return bbolt.ErrDatabaseNotOpen
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(id))
		if v == nil {
			return bbolt.ErrBucketNotFound
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) SaveWave(v domain.Wave) error           { return put(s, buckets[0], v.ID, v) }
func (s *Store) SaveTask(v domain.PickTask) error       { return put(s, buckets[1], v.ID, v) }
func (s *Store) SaveZone(v domain.Zone) error           { return put(s, buckets[2], v.ID, v) }
func (s *Store) SaveException(v domain.Exception) error { return put(s, buckets[3], v.ID, v) }
func (s *Store) Wave(id string) (domain.Wave, error) {
	var v domain.Wave
	e := get(s, buckets[0], id, &v)
	return v, e
}
func (s *Store) Task(id string) (domain.PickTask, error) {
	var v domain.PickTask
	e := get(s, buckets[1], id, &v)
	return v, e
}
func (s *Store) AllTasks() ([]domain.PickTask, error) {
	out := []domain.PickTask{}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, bbolt.ErrDatabaseNotOpen
	}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[1]).ForEach(func(_, v []byte) error {
			var t domain.PickTask
			if e := json.Unmarshal(v, &t); e != nil {
				return e
			}
			out = append(out, t)
			return nil
		})
	})
	return out, e
}
