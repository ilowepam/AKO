package model

type Company struct {
	Id          int
	Name        string
	CommanderId int
}

func NewCompanyObj(pName string, pCommanderId int) Company {
	return Company{Name: pName, CommanderId: pCommanderId}
}
