package models

//Company estructura de una empresa
type Company struct {
	ID              string      `json:"id"`
	ActivityCode    string      `json:"activityCode"`
	Address         string      `json:"address"`
	Activity        string      `json:"activity"`
	Auth            Auth        `json:"auth"`
	CpnCode         string      `json:"cpnCode"`
	CpnID           int         `json:"cpnId"`
	ContactEmail    string      `json:"contactEmail"`
	CpnName         string      `json:"cpnName"`
	City            string      `json:"city"`
	DailySummary    bool        `json:"dailySummary"`
	District        string      `json:"district"`
	GenerationDate  int         `json:"generationDate"`
	IsActive        bool        `json:"isActive"`
	InCertification bool        `json:"inCertification"`
	InProduction    bool        `json:"inProduction"`
	PrintedActivity string      `json:"printedActivity"`
	SenderCode      string      `json:"senderCode"`
	SiiUnit         string      `json:"siiUnit"`
	Certificate     Certificate `json:"certificate"`
}

//Maullin estructura del ambiente de certificacion
type Maullin struct {
	Date    int `json:"date"`
	Ticket  int `json:"ticket"`
	Invoice int `json:"invoice"`
}

//Palena estructura del ambiente de productivo
type Palena struct {
	TicketDate  interface{} `json:"ticketDate"`
	InvoiceDate interface{} `json:"invoiceDate"`
	Ticket      interface{} `json:"ticket"`
	Invoice     interface{} `json:"invoice"`
}

//Auth tipo de ambiente que se usara
type Auth struct {
	Maullin Maullin `json:"maullin"`
	Palena  Palena  `json:"palena"`
}

//Certificate struct of certificate
type Certificate struct {
	ID             string `json:"id"`
	ContentType    string `json:"contentType"`
	Base64         string `json:"base64"`
	ExpirationDate int    `json:"expirationDate"`
	GenerationDate int    `json:"generationDate"`
	Pass           string `json:"pass"`
	Revoked        bool   `json:"revoked"`
}

//Companies struct of a company from bsale
type Companies struct {
	CpnID   int    `json:"cpnId"`
	CpnCode string `json:"cpnCode"`
	CpnName string `json:"cpnName"`
}
