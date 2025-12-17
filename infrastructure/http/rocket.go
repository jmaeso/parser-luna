package http

import "github.com/jmaeso/parser-luna/domain"

// Rocket defines the entity for returning Rocket states.
type Rocket struct {
	ID              string `json:"channel"`
	Name            string `json:"type"`
	Mission         string `json:"mission"`
	Speed           int    `json:"speed"`
	ExplosionReason string `json:"explosionReason,omitempty"`
}

func newRocketFromDomain(r domain.Rocket) Rocket {
	return Rocket{
		ID:              r.ID,
		Name:            r.Name,
		Mission:         r.Mission,
		Speed:           r.Speed,
		ExplosionReason: r.ExplosionReason,
	}
}
