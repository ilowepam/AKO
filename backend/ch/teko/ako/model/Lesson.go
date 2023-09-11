package model

type Lesson struct {
	Id   int
	Name string
	Time int
	Desc string
}

func New(pName string, pTime int, pDesc string) Lesson {
	return Lesson{Name: pName, Time: pTime, Desc: pDesc}
}
