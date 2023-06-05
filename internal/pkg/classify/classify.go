package classify

type WebInfo struct {
	Title               string
	Content             string
	StatusCode          int
	Keywords            string
	Description         string
	Language            string
	ExternLinks         []string
	NumberOfExternLinks int
	IsSearch            int
	IsLogin             int
	IsRegister          bool
}

type WhoisInfo struct {
	RegisterName       string
	Registrar          string
	RegistrantPhone    string
	RegistrantEmail    string
	RegistrantOrg      string
	RegistrantCountry  string
	RegistrantCity     string
	NameServers        string
	RegisterTime       string
	RegisterExpireTime string
	EnterPrisePerson   string
	EnterPriseName     string
}

type WebSiteSeo struct {
	AlexaRanking      int
	ComprehensiveRank int
	WebuvMax          int
	WebuvMin          int
}

type Wappalyzer struct {
	WebServers          string
	WebFrameworks       string
	ProgrammingLanguage string
	CmsType             string
	SiteDatabase        string
	OperatingSystem     string
	OtherArchitecture   string
}
