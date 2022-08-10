package entities_core

type Address struct {
	Id           int64  `json:"id"`
	Public_area  string `json:"public_area"`
	Street       string `json:"street"`
	Number       int64  `json:"number"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}
