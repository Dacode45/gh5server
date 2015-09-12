package main

//Enum that switches between the municipal court types
type MunicipalCourtEnum int

const (
	NO_MUNICIPAL_COURT        MunicipalCourtEnum = 0
	HAS_MUNICIPAL_COURT       MunicipalCourtEnum = 1
	NO_MUNICIPAL_COURT_STLCO  MunicipalCourtEnum = 2
	FRONTENAC_MUNICIPAL_COURT MunicipalCourtEnum = 3
	NORMANDY_MUNICIPAL_COURT  MunicipalCourtEnum = 4
)

func (m MunicipalCourtEnum) String() string {
	s := ""
	switch m {
	case HAS_MUNICIPAL_COURT:
		s += "Has a Municipal Court"
	case NO_MUNICIPAL_COURT:
		s += "No Municipal Court"
	case NO_MUNICIPAL_COURT_STLCO:
		s += "No Municipal Court (stlco)"
	case FRONTENAC_MUNICIPAL_COURT:
		s += "Has a Frontenac Municiapl Court"
	case NORMANDY_MUNICIPAL_COURT:
		s += "Has a Normandy Municiapl Court"
	}
	return s
}

type Municipality struct {
	Id                    int                `json:"id"`
	Name                  string             `json:"name"`
	MunicipalCourt        MunicipalCourtEnum `json:"municipal_court"`
	MunicipalWebsite      string             `json:"municiapal_website"`
	MunicipalCourtWebsite string             `json:"municipal_court_website"`
	CourtDocketListed     bool               `json:"court_docket_listed"`
	CourtClerkPhoneListed bool               `json:"court_clerk_phone_listed"`
	CourtClerkPhone       string             `json:"court_clerk_phone"`
	InfoSource            string             `json:"info_source"`
	OnlinePaymentProvider string             `json:"online_payment_provider"`
	HasDressCode          bool               `json:"has_dress_code"`
	FineScheduleListed    bool
	MunicipalCodes        string `json:"municipal_codes"`
	OrderSeen             string `json:"order_seen"`
	DirectionsListed      bool   `json:"directions_listed"`
}

type Municipalities []Municipality
