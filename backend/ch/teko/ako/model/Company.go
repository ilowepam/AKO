package model

type Company struct {
	Id           int
	Headquarters string
	Commander    Soldier
	Platoons     []Platoon
}
