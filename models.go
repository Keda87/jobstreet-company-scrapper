package main

type Payload struct {
	URL       string
	UserAgent string
}

type Company struct {
	Address      string `json:"address"`
	CompanyName  string `json:"company_name"`
	CompanySize  string `json:"company_size"`
	Industry     string `json:"industry"`
	MapAddress   string `json:"map_address"`
	MapLatitude  string `json:"map_latitude"`
	MapLongitude string `json:"map_longitude"`
}

type Response struct {
	Data []Company `json:"data"`
}
