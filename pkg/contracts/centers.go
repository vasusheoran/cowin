package contracts

type Center struct {
	CenterID      int          `json:"center_id,omitempty"`
	Name          string       `json:"name,omitempty"`
	NameL         string       `json:"name_l,omitempty"`
	Address       string       `json:"address,omitempty"`
	AddressL      string       `json:"address_l,omitempty"`
	StateName     string       `json:"state_name,omitempty"`
	StateNameL    string       `json:"state_name_l,omitempty"`
	DistrictName  string       `json:"district_name,omitempty"`
	DistrictNameL string       `json:"district_name_l,omitempty"`
	BlockName     string       `json:"block_name,omitempty"`
	BlockNameL    string       `json:"block_name_l,omitempty"`
	Pincode       int          `json:"pincode,omitempty"`
	Lat           float64      `json:"lat,omitempty"`
	Long          float64      `json:"long,omitempty"`
	From          string       `json:"from,omitempty"`
	To            string       `json:"to,omitempty"`
	FeeType       string       `json:"fee_type,omitempty"`
	VaccineFees   []VaccineFee `json:"vaccine_fees,omitempty"`
	Sessions      []Session    `json:"sessions,omitempty"`
}
