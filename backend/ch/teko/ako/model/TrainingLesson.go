package model

type TrainingLesson struct {
	id   int
	name string
	time int
	desc string
}

func CreateTrainingLesson(pName string, pTime int, pDesc string) TrainingLesson {
	return TrainingLesson{name: pName, time: pTime, desc: pDesc}
}
