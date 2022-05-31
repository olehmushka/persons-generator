package entities

type Health struct {
	Mental *MentalHealth
	Body   *BodyHealth
}

type MentalHealth struct{}

type BodyHealth struct{}
