package storage

import (
	"errors"

	"github.com/jmaeso/parser-luna/domain"
)

// ErrRocketNotFound is an error to be used when a specific rocket data was expected but not found in the storage.
var ErrRocketNotFound = errors.New("rocket not found")

// The Messages interface defines the methods any storage mechanism needs to implement to interact with Messages.
type Messages interface {
	// Insert is expected to be a fast operation.
	// Can return an error since storage access is not always guaranteed.
	Insert(m domain.Message) error

	// GetSortedByRocketID expects all the messages to be sorted by EventNumber.
	// Can return an error since storage access is not always guaranteed.
	GetSortedByRocketID(id string) ([]domain.Message, error)

	// GetAllSorted expects all the messages for each rocket to be sorted by EventNumber.
	// Can return an error since storage access is not always guaranteed.
	GetAllSorted() ([][]domain.Message, error)
}
