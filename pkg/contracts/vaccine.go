package contracts

type VaccineFee struct {
	Vaccine string `json:"vaccine,omitempty"`
	Fee     string `json:"fee,omitempty"`
}
