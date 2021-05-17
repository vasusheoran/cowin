package contracts

type Filter struct {
	ID       string `json:"id,omitempty"`
	Location int    `json:"location,omitempty"`
	DoseType int    `json:"dose_type,omitempty"`
	Age      int    `json:"age,omitempty"`
	Vaccine  string `json:"vaccine,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
}

type Notify map[string]Filter

type FiterResponse struct {
	Name     string
	Address  string
	Pin      int
	CenterID int
	Filter   Filter
	Session  Session
}
