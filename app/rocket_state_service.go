package app

import (
	"fmt"

	"github.com/jmaeso/parser-luna/domain"
	"github.com/jmaeso/parser-luna/infrastructure/storage"
)

// RocketStateService is the application component responsible for computing the state
// of rockets by replaying the messages/events from the database.
// It expects the events to come sorted by order in order to go through them.
type RocketStateService struct {
	messagesStore storage.Messages
}

// NewRocketStateService creates a new RocketStateService with the given message repository
func NewRocketStateService(messageRepo storage.Messages) RocketStateService {
	return RocketStateService{
		messagesStore: messageRepo,
	}
}

// BuildAllRocketsState goes through all the messages in the system and builds the state
// for all the rockets present in those messages.
func (s *RocketStateService) BuildAllRocketsState() ([]domain.Rocket, error) {
	allRocketMessages, err := s.messagesStore.GetAllSorted()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving messages for all rockets Err: %w", err)
	}

	var rockets []domain.Rocket

	for _, rocketMessages := range allRocketMessages {
		rocket := s.buildRocketState(rocketMessages)
		rockets = append(rockets, rocket)
	}

	return rockets, nil
}

// BuildRocketState retrieves all messages for the given rocket ID and computes its final state.
// Can return storage.ErrRocketNotFound if the rocket is not found.
func (s *RocketStateService) BuildRocketState(rocketID string) (*domain.Rocket, error) {
	messages, err := s.messagesStore.GetSortedByRocketID(rocketID)
	if err != nil {
		return nil, fmt.Errorf("failed retrieving messages for single rocket Err: %w", err)
	}

	rocket := s.buildRocketState(messages)

	return &rocket, nil
}

func (s *RocketStateService) buildRocketState(messages []domain.Message) domain.Rocket {
	var rocket domain.Rocket

	for _, msg := range messages {
		if rocket.ID == "" {
			rocket.ID = msg.Metadata.Channel
		}

		if msg.Event.RocketLaunched != nil {
			rocket.Name = msg.Event.RocketLaunched.Type
			rocket.Speed = msg.Event.RocketLaunched.LaunchSpeed
			rocket.Mission = msg.Event.RocketLaunched.Mission
		}

		if msg.Event.RocketSpeedIncreased != nil {
			rocket.Speed += msg.Event.RocketSpeedIncreased.By
		}

		if msg.Event.RocketSpeedDecreased != nil {
			rocket.Speed -= msg.Event.RocketSpeedDecreased.By
		}

		if msg.Event.RocketExploded != nil {
			rocket.ExplosionReason = msg.Event.RocketExploded.Reason
		}

		if msg.Event.RocketMissionChanged != nil {
			rocket.Mission = msg.Event.RocketMissionChanged.NewMission
		}
	}

	return rocket
}
