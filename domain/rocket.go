package domain

// Rocket is the system's definition of a Rocket and the attributes which define its state.
type Rocket struct {
	ID              string
	Name            string
	Mission         string
	Speed           int
	ExplosionReason string // Empty if no explosion.
}
