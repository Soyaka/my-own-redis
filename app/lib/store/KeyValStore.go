package store

// import "sync"

type Storage struct {
	// mu    *sync.RWMutex
	Store map[string]string
}

type StoreWorker interface {
	SetValue(key, value string)
	GetValue(key string) (string, bool)
}

func NewStorage() *Storage {
	return &Storage{
		Store: make(map[string]string),
	}
}

func (s *Storage) GetValue(key string) (string, bool) {
	//s.mu.RLock()
	value, OK := s.Store[key]
	//s.mu.RUnlock()
	if !OK {
		return "", OK
	}
	return value, OK
}

func (s *Storage) SetValue(key, value string) {
	//s.mu.Lock()
	s.Store[key] = value
	//s.mu.Unlock()
}
