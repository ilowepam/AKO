package model

type Platoon struct {
	id        int
	name      string
	strength  int
	location  string
	commander AdA
	soldiers  []AdA
}
