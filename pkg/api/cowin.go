package api

import (
	"time"
)

type Cowin interface {
	GetCalendarByDistrict(district int, date time.Time) error
	Schedule(cronInterval string)
	AddFilter(districts []int, doseType, age int, vaccine string)
}
