package notify

import (
	"context"
	"cowin/pkg/api"
	"cowin/pkg/constants"
	"cowin/pkg/contracts"
	"cowin/utils"
	"fmt"
	"os"
	"strconv"

	"github.com/gen2brain/beeep"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/robfig/cron/v3"
)

var (
	cronJob *cron.Cron
)

type notifyService struct {
	logger    log.Logger
	filterMap map[int]contracts.Notify
}

func New(logger log.Logger) api.Notify {
	return &notifyService{
		logger:    logger,
		filterMap: map[int]contracts.Notify{},
	}
}

func (ns *notifyService) Add(filter contracts.Filter) error {
	val, _ := ns.filterMap[filter.Location]
	if val == nil {
		ns.filterMap[filter.Location] = contracts.Notify{}
	}

	ns.filterMap[filter.Location][filter.ID] = filter
	return nil
}

func (ns *notifyService) Remove(filterID string, district int) error {
	if _, ok := ns.filterMap[district]; !ok {
		return nil
	}
	delete(ns.filterMap[district], filterID)
	return nil
}

func (ns *notifyService) ScheduleNotifications(cronInterval string) error {
	notifyCron := fmt.Sprintf(constants.CronJob, cronInterval)
	var ctx context.Context

	if cronJob != nil {
		level.Debug(ns.logger).Log("msg", "Stopping existing cron job.")
		ctx = cronJob.Stop()
		<-ctx.Done()
		level.Debug(ns.logger).Log("msg", "Creating new schedule.")
	} else {
		cronJob = cron.New()
	}

	level.Info(ns.logger).Log("Info", "Parsed cron URL", "ScheduleNotifications", notifyCron)
	cronJob.AddFunc(notifyCron, func() {
		ns.filterAndNotify()
	})

	cronJob.Start()
	return nil
}

func (ns *notifyService) filterAndNotify() error {
	files, err := os.ReadDir(constants.DistrictDirPath)
	if err != nil {
		return err
	}

	for _, f := range files {
		var data contracts.GetCalendarByDistrictResponse
		district, err := utils.GetDistrictFromFileName(f.Name())
		if err != nil {
			level.Error(ns.logger).Log("msg", "Unable to parse district from filename", "file", f)
			continue
		}

		path := fmt.Sprintf(constants.FilePath, constants.DistrictDirPath, strconv.Itoa(district))
		err = utils.Load(path, &data)
		if err != nil {
			level.Error(ns.logger).Log("msg", "Unable to read data", "file", f)
			continue
		}

		ns.Notify(data.Centers, district)
	}
	return nil
}

func (ns *notifyService) Notify(centers []contracts.Center, district int) {
	notify := ns.filterMap[district]
	for _, center := range centers {
		for _, session := range center.Sessions {
			for _, value := range notify {
				if ns.validateSession(session, value.Age, value.DoseType, value.Vaccine) {
					fr := contracts.FiterResponse{
						Name:     center.Name,
						Address:  center.Address,
						Pin:      center.Pincode,
						CenterID: center.CenterID,
						Filter:   value,
						Session:  session,
					}
					ns.sendAlert(fr)
				}
			}
		}
	}
}

func (ns *notifyService) sendAlert(filterResponse contracts.FiterResponse) {
	if filterResponse.Session.AvailableCapacityDose2 < constants.MinimumAlertVal {
		return
	}
	if filterResponse.Session.AvailableCapacityDose1 < constants.MinimumAlertVal {
		return
	}

	title := fmt.Sprintf(constants.TitleTemplate, filterResponse.Pin, filterResponse.Filter.Age, filterResponse.Name)
	responseBody := fmt.Sprintf(constants.SubtitleTemplate,
		filterResponse.Address,
		filterResponse.Session.AvailableCapacityDose1,
		filterResponse.Session.AvailableCapacityDose2,
		filterResponse.Session.Date)

	level.Info(ns.logger).Log("msg", "Valid session found", "title", title, "responseBody", responseBody)
	err := beeep.Alert(title, responseBody, "")
	if err != nil {
		level.Error(ns.logger).Log("msg", "Unable to beeep", "err", err)
	}
}

func (ns *notifyService) validateSession(session contracts.Session, age int, doseType int, vaccine string) bool {
	if len(vaccine) != 0 && vaccine != session.Vaccine {
		return false
	}

	if age < session.MinAgeLimit {
		return false
	}

	switch doseType {
	case 1:
		if session.AvailableCapacityDose1 <= 0 {
			return false
		}
		break
	case 2:
		if session.AvailableCapacityDose2 <= 0 {
			return false
		}
		break
	default:
		level.Error(ns.logger).Log("msg", "Dose type not found", "type", doseType)
	}

	return true
}
