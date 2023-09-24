package model

type Rank string

const (
	General               string = "General"
	CorpsCommander        string = "Korpskommandant"
	DivisionOfficer       string = "Divisionär"
	Brigadier             string = "Brigadier"
	SpecialistOfficer     string = "Fachoffizier"
	Colonel               string = "Oberst"
	LieutenantColonel     string = "Oberstleutnant"
	Major                 string = "Major"
	Captain               string = "Hauptmann"
	FirstLieutenant       string = "Oberleutnant"
	SecondLieutenant      string = "Leutnant"
	ChiefAdjutant         string = "Chefadjutant"
	HeadAdjutant          string = "Hauptadjutant"
	StaffAdjutant         string = "Stabsadjutant"
	AdjutantSergeantMajor string = "Adjutant Unteroffizier"
	SergeantMajor         string = "Hauptfeldweibel"
	Fourier               string = "Fourier"
	WarrantOfficer        string = "Feldweibel"
	SeniorSergeant        string = "Oberwachtmeister"
	Sergeant              string = "Wachtmeister"
	LanceCorporal         string = "Korporal"
	Corporal              string = "Obergefreiter"
	Private               string = "Gefreiter"
	Soldier               string = "Soldat"
)

var AllRanks = []string{
	General,
	CorpsCommander,
	DivisionOfficer,
	Brigadier,
	SpecialistOfficer,
	Colonel,
	LieutenantColonel,
	Major,
	Captain,
	FirstLieutenant,
	SecondLieutenant,
	ChiefAdjutant,
	HeadAdjutant,
	StaffAdjutant,
	AdjutantSergeantMajor,
	SergeantMajor,
	Fourier,
	WarrantOfficer,
	SeniorSergeant,
	Sergeant,
	LanceCorporal,
	Corporal,
	Private,
	Soldier,
}
