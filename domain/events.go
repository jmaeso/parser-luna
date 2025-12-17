package domain

import "time"

// Message represents the complete message structure received from rockets constantly/periodically.
type Message struct {
	Metadata Metadata
	Event    EventData
}

// Metadata contains the message metadata.
type Metadata struct {
	Channel     string // Seems it could be UUID, but better not asume it.
	EventNumber int
	EventTime   time.Time
	EventType   EventType
}

// EventType is defined to control which events are known by the system.
type EventType string

const (
	EventTypeRocketLaunched       EventType = "RocketLaunched"
	EventTypeRocketSpeedIncreased EventType = "RocketSpeedIncreased"
	EventTypeRocketSpeedDecreased EventType = "RocketSpeedDecreased"
	EventTypeRocketExploded       EventType = "RocketExploded"
	EventTypeRocketMissionChanged EventType = "RocketMissionChanged"
)

// EventData defines all the possible events inside the same structure for easy definition and access.
type EventData struct {
	RocketLaunched       *RocketLaunched
	RocketSpeedIncreased *RocketSpeedIncreased
	RocketSpeedDecreased *RocketSpeedDecreased
	RocketExploded       *RocketExploded
	RocketMissionChanged *RocketMissionChanged
}

// RocketLaunched defines a rocket launch event.
type RocketLaunched struct {
	Type        string
	LaunchSpeed int
	Mission     string
}

// RocketSpeedIncreased defines a speed increase event.
type RocketSpeedIncreased struct {
	By int
}

// RocketSpeedDecreased defines a speed decrease event.
type RocketSpeedDecreased struct {
	By int
}

// RocketExploded defines an explosion event.
type RocketExploded struct {
	Reason string
}

// RocketMissionChanged defines a mission change event.
type RocketMissionChanged struct {
	NewMission string
}
