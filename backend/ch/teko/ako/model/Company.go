package model

type Company struct {
	id           int
	headquarters string
	commander    AdA
	platoons     []Platoon
}
