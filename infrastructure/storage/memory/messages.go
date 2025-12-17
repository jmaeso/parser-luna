package memory

import (
	"sort"
	"sync"

	"github.com/jmaeso/parser-luna/domain"
	"github.com/jmaeso/parser-luna/infrastructure/storage"
)

// MessagesStore is the in-memory implementation of storage.Messages.
// Internally stores all the events for the same message channel by order of reception.
//
// For the sake of simplicity and speed, it uses the domain.Message type, ideally should define it's own Message type.
type MessagesStore struct {
	messages map[string][]domain.Message
	mu       sync.Mutex
}

// NewMessagesStore MUST be used to initialize the store to avoid nil pointer accessors.
func NewMessagesStore() *MessagesStore {
	return &MessagesStore{
		messages: make(map[string][]domain.Message),
	}
}

// Insert stores the new message without any sorting logic.
// The implementation is safe for concurrent usage.
func (s *MessagesStore) Insert(m domain.Message) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	channel := m.Metadata.Channel

	s.messages[channel] = append(s.messages[channel], m)

	return nil
}

// GetByRocketID returns all the messages for the given rocket ID.
// The returned messages are sorted by it's event number.
// Can return storage.ErrRocketNotFound if there are no messages for the given rocketID.
func (s *MessagesStore) GetSortedByRocketID(id string) ([]domain.Message, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	messages, ok := s.messages[id]
	if !ok {
		return nil, storage.ErrRocketNotFound
	}

	// Make a copy so we don't mutate the original slice
	sorted := make([]domain.Message, len(messages))
	copy(sorted, messages)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Metadata.EventNumber < sorted[j].Metadata.EventNumber
	})

	return sorted, nil
}

// GetAllSorted returns all the messages for all the (known) rockets.
// The returned messages are sorted by it's event number.
func (s *MessagesStore) GetAllSorted() ([][]domain.Message, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var allMessages [][]domain.Message

	for _, messages := range s.messages {
		// Make a copy so we don't mutate the original slice
		sorted := make([]domain.Message, len(messages))
		copy(sorted, messages)

		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].Metadata.EventNumber < sorted[j].Metadata.EventNumber
		})

		allMessages = append(allMessages, sorted)
	}

	return allMessages, nil
}
