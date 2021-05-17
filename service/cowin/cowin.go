package cowin

import (
	"context"
	err2 "cowin/err"
	"cowin/pkg/api"
	"cowin/pkg/constants"
	"cowin/pkg/contracts"
	"cowin/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

type cowinService struct {
	logger    log.Logger
	client    api.HTTPClient
	notify    api.Notify
	districts map[int]bool
	pincodes  map[int]bool
	cronJobs  []*cron.Cron
}

func New(logger log.Logger, notify api.Notify, client api.HTTPClient) api.Cowin {
	return &cowinService{
		logger:    logger,
		client:    client,
		notify:    notify,
		districts: map[int]bool{},
		pincodes:  map[int]bool{},
		cronJobs:  []*cron.Cron{},
	}
}

func (cs *cowinService) AddFilter(districts, pincodes []int, doseType, age int, vaccine string) {
	for _, district := range districts {
		uid := uuid.New().String()

		cs.notify.Add(contracts.Filter{
			ID:       uid,
			Location: district,
			DoseType: doseType,
			Age:      age,
			Vaccine:  vaccine,
		})

		cs.districts[district] = false
	}

	for _, pincode := range pincodes {
		uid := uuid.New().String()

		cs.notify.Add(contracts.Filter{
			ID:       uid,
			Location: pincode,
			DoseType: doseType,
			Age:      age,
			Vaccine:  vaccine,
		})

		cs.pincodes[pincode] = false
	}
}

var cronJob *cron.Cron

func (cs *cowinService) Schedule(cronInterval string) {

	if cronJob != nil {
		var ctx context.Context
		level.Debug(cs.logger).Log("msg", "Stopping existing cron job.")
		ctx = cronJob.Stop()
		<-ctx.Done()

	}

	level.Debug(cs.logger).Log("msg", "Creating new cron.")
	cronJob = cron.New()

	utils.MakeDirectoryIfNotExists(constants.DistrictDirPath, constants.PincodeDirPath)

	for key, _ := range cs.districts {
		district := key

		getCalendarByDistrictCron := fmt.Sprintf(constants.CronJob, cronInterval)
		level.Info(cs.logger).Log("Info", "Parsed cron URL", "GetCalendarByDistrictCron", getCalendarByDistrictCron, "district", district)
		cronJob.AddFunc(getCalendarByDistrictCron, func() {
			cs.GetCalendarByDistrict(district, time.Now())
		})

		// Add the cron job to the list
		cs.districts[district] = true
	}

	for key, _ := range cs.pincodes {
		pincode := key

		getCalendarByDistrictCron := fmt.Sprintf(constants.CronJob, cronInterval)
		level.Info(cs.logger).Log("Info", "Parsed cron URL", "GetCalendarByPinCron", getCalendarByDistrictCron, "pincode", pincode)
		cronJob.AddFunc(getCalendarByDistrictCron, func() {
			cs.GetCalendarByPincode(pincode, time.Now())
		})

		// Add the cron job to the list
		cs.districts[pincode] = true
	}

	cronJob.Start()
}

func (cs *cowinService) GetCalendarByDistrict(district int, date time.Time) error {
	placeholders := map[string]interface{}{
		constants.KeyBaseURI:    constants.ProductionServer,
		constants.KeyDistrictID: district,
		constants.KeyDate:       date.Format(constants.Layout),
	}

	uri, err := utils.ParseTemplateString(constants.CalendarByDistrict, placeholders)
	if err != nil {
		level.Error(cs.logger).Log("msg", "failed to parse template string url")
		return err
	}

	res, err := cs.invokeAPI(uri)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(constants.FilePath, constants.DistrictDirPath, strconv.Itoa(district))
	if err := utils.Save(path, res); err != nil {
		level.Error(cs.logger).Log("err", err)
		return err
	}

	level.Info(cs.logger).Log("msg", "Successfully updated data. Sending alerts if any ...", "district", district)
	cs.notify.Notify(res.Centers, district)

	return nil
}

func (cs *cowinService) GetCalendarByPincode(pincode int, date time.Time) error {
	placeholders := map[string]interface{}{
		constants.KeyBaseURI: constants.ProductionServer,
		constants.KeyPincode: pincode,
		constants.KeyDate:    date.Format(constants.Layout),
	}

	uri, err := utils.ParseTemplateString(constants.CalendarByPincode, placeholders)
	if err != nil {
		level.Error(cs.logger).Log("msg", "failed to parse template string url")
		return err
	}

	res, err := cs.invokeAPI(uri)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(constants.FilePath, constants.PincodeDirPath, strconv.Itoa(pincode))
	if err := utils.Save(path, res); err != nil {
		level.Error(cs.logger).Log("err", err)
		return err
	}

	level.Info(cs.logger).Log("msg", "Successfully updated data. Sending alerts if any ...", "pincode", pincode)
	cs.notify.Notify(res.Centers, pincode)

	return nil
}

func (cs *cowinService) invokeAPI(uri string) (*contracts.GetCalendarByDistrictResponse, error) {
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		level.Error(cs.logger).Log("msg", "Failed to form http request", "err", err.Error())
		return nil, err
	}

	req.Header.Add(constants.UserAgentHeaderyKey, constants.UserAgentHeaderValue)
	resp, err := cs.client.Do(req)
	if err != nil {
		level.Error(cs.logger).Log("err", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		if err != nil {
			level.Error(cs.logger).Log("err", err.Error(), "status", resp.StatusCode)
			return nil, err
		}
		responseMessage := string(respBytes)
		level.Error(cs.logger).Log("response", responseMessage, "status", resp.StatusCode)
		return nil, err2.FailedToMakeHTTPRequest
	}

	//level.Info(cs.logger).Log("resp", string(respBytes))

	centers := contracts.GetCalendarByDistrictResponse{}
	err = json.Unmarshal(respBytes, &centers)
	if err != nil {
		level.Error(cs.logger).Log("err", err, "msg", "Failed to parse json")
		return nil, err2.FailedToParseJson
	}

	return &centers, nil
}
