package contracts

type Session struct {
	SessionID              string   `json:"session_id,omitempty"`
	Date                   string   `json:"date,omitempty"`
	AvailableCapacity      int      `json:"available_capacity,omitempty"`
	AvailableCapacityDose1 int      `json:"available_capacity_dose1,omitempty"`
	AvailableCapacityDose2 int      `json:"available_capacity_dose2,omitempty"`
	MinAgeLimit            int      `json:"min_age_limit,omitempty"`
	Vaccine                string   `json:"vaccine,omitempty"`
	Slots                  []string `json:"slots,omitempty"`
}
