package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jmaeso/parser-luna/domain"
)

// PostMessagePayload defines the body payload received in POST /messages.
type PostMessagePayload struct {
	Metadata Metadata        `json:"metadata"`
	Message  json.RawMessage `json:"message"` // Can be any of the specific message types
}

// Metadata contains the message metadata.
type Metadata struct {
	Channel       string      `json:"channel"`
	MessageNumber int         `json:"messageNumber"`
	MessageTime   time.Time   `json:"messageTime"`
	MessageType   MessageType `json:"messageType"`
}

// MessageType is defined to easily list which events are expected to be received.
type MessageType string

const (
	MessageTypeRocketLaunched       MessageType = "RocketLaunched"
	MessageTypeRocketSpeedIncreased MessageType = "RocketSpeedIncreased"
	MessageTypeRocketSpeedDecreased MessageType = "RocketSpeedDecreased"
	MessageTypeRocketExploded       MessageType = "RocketExploded"
	MessageTypeRocketMissionChanged MessageType = "RocketMissionChanged"
)

// RocketLaunched defines a rocket launch event payload
type RocketLaunched struct {
	Type        string `json:"type"`
	LaunchSpeed int    `json:"launchSpeed"`
	Mission     string `json:"mission"`
}

// RocketSpeedIncreased defines a speed increase event payload
type RocketSpeedIncreased struct {
	By int `json:"by"`
}

// RocketSpeedDecreased defines a speed decrease event payload
type RocketSpeedDecreased struct {
	By int `json:"by"`
}

// RocketExploded defines an explosion event payload
type RocketExploded struct {
	Reason string `json:"reason"`
}

// RocketMissionChanged defines a mission change event payload
type RocketMissionChanged struct {
	NewMission string `json:"newMission"`
}

func (p PostMessagePayload) ToDomainMessage() (domain.Message, error) {
	metadata := domain.Metadata{
		Channel:     p.Metadata.Channel,
		EventNumber: p.Metadata.MessageNumber,
		EventTime:   p.Metadata.MessageTime,
		EventType:   domain.EventType(p.Metadata.MessageType),
	}

	var event domain.EventData

	switch p.Metadata.MessageType {
	case MessageTypeRocketLaunched:
		var msg RocketLaunched
		if err := json.Unmarshal(p.Message, &msg); err != nil {
			return domain.Message{}, fmt.Errorf("failed to unmarshal RocketLaunched: %w", err)
		}
		event.RocketLaunched = &domain.RocketLaunched{
			Type:        msg.Type,
			LaunchSpeed: msg.LaunchSpeed,
			Mission:     msg.Mission,
		}
	case MessageTypeRocketSpeedIncreased:
		var msg RocketSpeedIncreased
		if err := json.Unmarshal(p.Message, &msg); err != nil {
			return domain.Message{}, fmt.Errorf("failed to unmarshal RocketSpeedIncreased: %w", err)
		}
		event.RocketSpeedIncreased = &domain.RocketSpeedIncreased{
			By: msg.By,
		}
	case MessageTypeRocketSpeedDecreased:
		var msg RocketSpeedDecreased
		if err := json.Unmarshal(p.Message, &msg); err != nil {
			return domain.Message{}, fmt.Errorf("failed to unmarshal RocketSpeedDecreased: %w", err)
		}
		event.RocketSpeedDecreased = &domain.RocketSpeedDecreased{
			By: msg.By,
		}
	case MessageTypeRocketExploded:
		var msg RocketExploded
		if err := json.Unmarshal(p.Message, &msg); err != nil {
			return domain.Message{}, fmt.Errorf("failed to unmarshal RocketExploded: %w", err)
		}
		event.RocketExploded = &domain.RocketExploded{
			Reason: msg.Reason,
		}
	case MessageTypeRocketMissionChanged:
		var msg RocketMissionChanged
		if err := json.Unmarshal(p.Message, &msg); err != nil {
			return domain.Message{}, fmt.Errorf("failed to unmarshal RocketMissionChanged: %w", err)
		}
		event.RocketMissionChanged = &domain.RocketMissionChanged{
			NewMission: msg.NewMission,
		}
	default:
		return domain.Message{}, fmt.Errorf("unknown message type: %s", p.Metadata.MessageType)
	}

	return domain.Message{
		Metadata: metadata,
		Event:    event,
	}, nil
}
