package constants

var MinimumAlertVal = 5

const (
	ProductionServer   = "https://cdn-api.co-vin.in/api"
	CalendarByDistrict = "{{ .BaseURI }}/v2/appointment/sessions/calendarByDistrict?district_id={{ .DistrictID }}&date={{ .Date }}"
	CalendarByPincode  = "{{ .BaseURI }}/v2/appointment/sessions/calendarByPin?pincode={{ .Pincode }}&date={{ .Date }}"
	Layout             = "02-01-2006"
	DistrictDirPath    = "./output/districts"
	PincodeDirPath     = "./output/pincodes"
	FilePath           = "%s/%s.json"
	AlertsDirPath      = "./output/alerts"
	AlertsFilePath     = "%s/%s.json"
	TitleTemplate      = "[%d] [%d+] %s"
	SubtitleTemplate   = "%s, Dose1: %d, Dose2: %d, %s"
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
