package model

type Lesson struct {
	Id       int
	Name     string
	Duration int
	Desc     string
}

func NewLessonObj(pName string, pDuration int, pDesc string) Lesson {
	return Lesson{Name: pName, Duration: pDuration, Desc: pDesc}
}
