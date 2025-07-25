package calculator

import (
	"encoding/json"
	"fmt"
	"os"
)

const SaveFilePath = "./results.json"

type JSONStore struct {
	*ResultStore
}

func NewJSONStore() (*JSONStore, error) {
	store := &JSONStore{
		ResultStore: NewResultStore(),
	}

	if err := store.Load(); err != nil {
		return nil, fmt.Errorf("failed to load storage: %w", err)
	}

	return store, nil
}

func (s *JSONStore) Load() error {
	data, err := os.ReadFile(SaveFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // file doesn't exist, start with empty storage
		}
		return fmt.Errorf("failed to read storage file: %w", err)
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	return json.Unmarshal(data, &s.results)
}

func (s *JSONStore) Save() error {
	s.mu.RLock()
	data, err := json.Marshal(s.results)
	s.mu.RUnlock()

	if err != nil {
		return fmt.Errorf("failed to marshal storage: %w", err)
	}

	if err := os.WriteFile(SaveFilePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write storage file: %w", err)
	}

	return nil
}
