package store

import "time"

// import "sync"

type Storage struct {
	// mu    *sync.RWMutex
	Store map[string]Data
}

type Data struct {
	Value      string
	ExpriredAt time.Time
}
func NewData(value string, expire time.Time )*Data{
	return &Data{
		Value: value,
		ExpriredAt: expire,
	}
}
type StoreWorker interface {
	SetValue(key string, value Data)
	GetValue(key string) (string, bool)
}

func NewStorage() *Storage {
	return &Storage{
		Store: make(map[string]Data),
	}
}

func (s *Storage) GetValue(key string) (string, bool) {
	//s.mu.RLock()
	data, OK := s.Store[key]
	//s.mu.RUnlock()
	if !OK || data.ExpriredAt.Before(time.Now()){
		return "", OK
	}
	return data.Value, OK
}

func (s *Storage) SetValue(key string, Data *Data) {
	//s.mu.Lock()
	s.Store[key] = *Data
	//s.mu.Unlock()
}
