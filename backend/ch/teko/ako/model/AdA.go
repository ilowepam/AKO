package model

type AdA struct {
	Id        int
	Name      string
	Rank      string
	CompanyId int
}

func GetAdAHeaderNames() []string {
	return []string{"Id", "Name", "Rank", "Company"}
}
