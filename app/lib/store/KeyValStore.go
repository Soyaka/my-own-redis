package store

import "time"

// import "sync"

type Storage struct {
	// mu    *sync.RWMutex
	Store map[string]Data
}

type Data struct {
	Value     string
	ExpireAt time.Time
}

type StoreWorker interface {
	SetValue(key, value string)
	GetValue(key string) (string, bool)
}

func NewStorage() *Storage {
	return &Storage{
		Store: make(map[string]Data),
	}
}

func (s *Storage) GetValue(key string) (string, bool) {
	//s.mu.RLock()
	// defer s.mu.RUnlock()

	data, OK := s.Store[key]
	if !OK || !time.Now().Before(data.ExpireAt) {
		return "", OK
	}
	return data.Value, OK
}

func (s *Storage) SetValue(key string, data Data) {
	//s.mu.Lock()
	//defer s.mu.Unlock()
	s.Store[key] = data
	
}
