package contracts

import "time"

type State struct {
	StateId    int8   `json:"state_id"`
	StateName  string `json:"state_name"`
	StateNameL string `json:"state_name_l"`
}

type District struct {
	StateId       int8   `json:"state_id"`
	DistrictId    int8   `json:"district_id"`
	DistrictName  string `json:"district_name"`
	DistrictNameL string `json:"district_name_l"`
}

type Cowin struct {
	State    State
	District District
	Date     time.Time
}
