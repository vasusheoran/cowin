package constants

import (
	"path/filepath"
)

var MinimumAlertVal = 5

const (
	ProductionServer   = "https://cdn-api.co-vin.in/api"
	CalendarByDistrict = "{{ .BaseURI }}/v2/appointment/sessions/calendarByDistrict?district_id={{ .DistrictID }}&date={{ .Date }}"
	CalendarByPincode  = "{{ .BaseURI }}/v2/appointment/sessions/calendarByPin?pincode={{ .Pincode }}&date={{ .Date }}"
	Layout             = "02-01-2006"
	TitleTemplate      = "[%d] [%d+] %s"
	SubtitleTemplate   = "%s, Dose1: %d, Dose2: %d, %s"
)

var (
	DistrictDirPath = filepath.Join("output", "districts")
	PincodeDirPath  = filepath.Join("output", "pincodes")
	AlertsDirPath   = filepath.Join("output", "alerts")
	FilePath        = filepath.Join("%s", "%s.json")
)

const (
	KeyBaseURI           = "BaseURI"
	KeyDistrictID        = "DistrictID"
	KeyPincode           = "Pincode"
	KeyDate              = "Date"
	UserAgentHeaderyKey  = "User-Agent"
	UserAgentHeaderValue = "Mozilla/5.0"

	COVAX      = "COVAXIN"
	COVISHIELD = "COVISHIELD"

	//https://apisetu.gov.in/public/marketplace/api/cowin/cowin-public-v2#/Appointment%20Availability%20APIs/calendarByCenter
	NWD    = 143
	SWD    = 150
	Rewari = 202

	CronJob = "@every %s"
)
