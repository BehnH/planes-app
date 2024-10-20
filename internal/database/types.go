package database

import "fmt"

type Aircraft struct {
	Id             string `json:"id,omitempty"`
	Manufacturer   string `json:"manufacturer"`
	Image          string `json:"image"`
	Overview       string `json:"overview"`
	EnteredService string `json:"entered_service"`
	Variant        string `json:"variant"`
	IcaoCode       string `json:"icao_code"`
}

func (a Aircraft) GetName() string {
	return fmt.Sprintf("%s %s", a.Manufacturer, a.Variant)
}

func (a Aircraft) GetFileUrl() string {
	return a.Image
}

func (a Aircraft) GetOverview() string {
	return a.Overview
}
