package api

import (
	"time"
)

type Cowin interface {
	GetCalendarByDistrict(district int, date time.Time) error
	Schedule(cronInterval string)
	AddFilter(districts, pincodes []int, doseType, age int, vaccine string)
}
