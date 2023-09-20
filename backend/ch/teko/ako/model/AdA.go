package model

type AdA struct {
	Id        int
	Name      string
	Rank      Rank
	CompanyId int
}

func NewAdAObj(pName string, pRank Rank, pCompanId int) AdA {
	return AdA{Name: pName, Rank: pRank, CompanyId: pCompanId}
}
