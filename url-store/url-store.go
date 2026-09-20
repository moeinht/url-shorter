package urlstore

import (
	"errors"
	"sync"
	"time"
	"url-shorter/hash"
)

type URLData struct {
	url        string
	createTime time.Time
}

type URLStore struct {
	mu   sync.RWMutex
	urls map[string]URLData
}

var (
	ErrNotFound = errors.New("url not found")
	ErrExpired  = errors.New("url expired")
)

func CreateStore() *URLStore {
	return &URLStore{
		urls: make(map[string]URLData),
	}
}

func (s *URLStore) CreateURL(original string) string {
	hashed := hash.GenerateCode()
	s.mu.Lock()
	defer s.mu.Unlock()

	for {
		if s.isExist(hashed) {
			hashed = hash.GenerateCode()
			continue
		}

		break
	}

	s.urls[hashed] = URLData{
		url:        original,
		createTime: time.Now(),
	}
	return hashed
}

func (s *URLStore) CleanupExpired() {
	s.mu.Lock()
	defer s.mu.Unlock()

	currentTime := time.Now()

	for key, value := range s.urls {
		if currentTime.After(value.createTime.Add(30 * 24 * time.Hour)) {
			delete(s.urls, key)
		}
	}
}

func (s *URLStore) GetURL(hash string) (URLData, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	hashData, exist := s.urls[hash]
	if !exist {
		return URLData{}, ErrNotFound
	}

	currentTime := time.Now()

	if currentTime.After(hashData.createTime.Add(30 * 24 * time.Hour)) {
		return URLData{}, ErrExpired
	}

	return hashData, nil
}

func (s *URLStore) isExist(hash string) bool {
	_, exists := s.urls[hash]

	return exists
}

func (s *URLStore) IsExist(hash string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.isExist(hash)
}
