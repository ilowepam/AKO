package model

type Platoon struct {
	Name      string
	Strength  int
	Location  string
	Commander Soldier
	Soldiers  []Soldier
}
