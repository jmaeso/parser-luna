package app

import (
	"testing"

	"github.com/jmaeso/parser-luna/domain"
)

func TestBuildRocketState(t *testing.T) {
	testCases := []struct {
		name                string
		inputMessages       []domain.Message
		expectedRocketState domain.Rocket
	}{
		{
			name: "Should add and substract speed properly",
			inputMessages: []domain.Message{
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketLaunched: &domain.RocketLaunched{Type: "Tintin Moon Rocket", LaunchSpeed: 100, Mission: "Moon landing"}},
				},
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketSpeedIncreased: &domain.RocketSpeedIncreased{By: 100}},
				},
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketSpeedDecreased: &domain.RocketSpeedDecreased{By: 10}},
				},
			},
			expectedRocketState: domain.Rocket{
				ID:      "123",
				Name:    "Tintin Moon Rocket",
				Mission: "Moon landing",
				Speed:   190,
			},
		},
		{
			name: "Should update mission name",
			inputMessages: []domain.Message{
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketLaunched: &domain.RocketLaunched{Type: "Tintin Moon Rocket", LaunchSpeed: 100, Mission: "Moon landing"}},
				},
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketMissionChanged: &domain.RocketMissionChanged{NewMission: "Asteroid exploration"}},
				},
			},
			expectedRocketState: domain.Rocket{
				ID:      "123",
				Name:    "Tintin Moon Rocket",
				Mission: "Asteroid exploration",
				Speed:   100,
			},
		},
		{
			name: "Should report the last known speed before crashing",
			inputMessages: []domain.Message{
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketLaunched: &domain.RocketLaunched{Type: "Tintin Moon Rocket", LaunchSpeed: 100, Mission: "Moon landing"}},
				},
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketSpeedIncreased: &domain.RocketSpeedIncreased{By: 100}},
				},
				{
					Metadata: domain.Metadata{Channel: "123"},
					Event:    domain.EventData{RocketExploded: &domain.RocketExploded{Reason: "PRESSURE_VESSEL_FAILURE"}},
				},
			},
			expectedRocketState: domain.Rocket{
				ID:              "123",
				Name:            "Tintin Moon Rocket",
				Mission:         "Moon landing",
				Speed:           200,
				ExplosionReason: "PRESSURE_VESSEL_FAILURE",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := NewRocketStateService(nil)

			rocket := s.buildRocketState(tc.inputMessages)

			if rocket != tc.expectedRocketState {
				t.Errorf("expected rocket state %+v, got %+v", tc.expectedRocketState, rocket)
			}
		})
	}
}
