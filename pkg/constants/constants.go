package constants

const (
	ProductionServer   = "https://cdn-api.co-vin.in/api"
	CalendarByDistrict = "{{ .BaseURI }}/v2/appointment/sessions/calendarByDistrict?district_id={{ .DistrictID }}&date={{ .Date }}"
	Layout             = "02-01-2006"
	DistrictDirPath    = "./output/districts"
	AlertsDirPath      = "./output/alerts"
	DistrictFilePath   = "%s/%s.json"
	AlertsFilePath     = "%s/%s.json"
	DistrictFileLayout = "02-01-2006T15:04"
)

const (
	KeyBaseURI           = "BaseURI"
	KeyDistrictID        = "DistrictID"
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
