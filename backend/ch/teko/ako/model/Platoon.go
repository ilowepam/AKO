package model

type Platoon struct {
	id        int
	name      string
	strength  int
	location  string
	commander AdA
	soldiers  []AdA
}

func CreatePlatoon(pName string, pTime int, pDesc string) TrainingLesson {
	return TrainingLesson{name: pName, time: pTime, desc: pDesc}
}
