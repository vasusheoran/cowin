package main

import (
	"cowin/pkg/constants"
	"cowin/pkg/contracts"
	"cowin/service/cowin"
	"cowin/service/notify"
	"crypto/tls"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

//https://cdn-api.co-vin.in/api/v2/admin/location/states
//https://cdn-api.co-vin.in/api/v2/admin/location/districts/9

const Interval = "5m"

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = log.WithPrefix(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	flg, err := parseFlag(logger)
	if err != nil {
		level.Error(logger).Log("err", err)
	}

	ch := make(chan bool)
	schedule(logger, flg)
	<-ch

}
func schedule(logger log.Logger, inputFlag *contracts.InputFlags) error {

	client := &http.Client{Timeout: 10 * time.Second, Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}

	ns := notify.New(logger)
	service := cowin.New(logger, ns, client)

	districts, err := getDistricts(inputFlag.Districts)
	if err != nil {
		return err
	}

	for _, val := range inputFlag.Filters {
		filter := strings.Split(val, ",")
		if len(filter) > 3 || len(filter) <= 0 {
			//TODO: move to errors
			return errors.New("invalid filter")
		}

		doseType, err := strconv.Atoi(filter[0])
		if err != nil || doseType < 0 || doseType > 2 {
			return errors.New("invalid dose type")
		}

		age, err := strconv.Atoi(filter[1])
		if err != nil || age < 0 || age > 100 {
			return errors.New("invalid age")
		}

		var vaccine string
		switch filter[2] {
		case "1":
			vaccine = constants.COVAX
			break
		case "2":
			vaccine = constants.COVISHIELD
			break
		default:
			vaccine = constants.COVAX
		}

		level.Info(logger).Log("msg", "Adding filter", "districts", fmt.Sprint(districts), "doseType", doseType, "age", age, "vaccine", vaccine)
		service.AddFilter(districts, doseType, age, vaccine)

	}
	service.Schedule(inputFlag.Interval)
	return nil
}

func getDistricts(districts string) ([]int, error) {
	a := strings.Split(districts, ",")
	b := []int{}
	for _, v := range a {

		val, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		b = append(b, val)
	}
	return b, nil
}

func parseFlag(logger log.Logger) (*contracts.InputFlags, error) {
	var inputFlags contracts.InputFlags

	flags := flag.NewFlagSet("cowin", flag.ExitOnError)
	flags.StringVar(&inputFlags.Districts, "districts", "", "<district>,<district>,... eg: 143,150 ")
	flags.Var(&inputFlags.Filters, "filter", "<dose type>,<minimum age>,<vaccine> eg: 1,18,1 [Dose Type: 1st dose (1), 2nd dose (2)] [Vaccine: COVAXIN (1), COVISHIELD (2)]")
	flags.StringVar(&inputFlags.Interval, "interval", "5m", "Interval [1s/2m/3h]")
	flags.BoolVar(&inputFlags.Help, "help", false, "Set to true for printing usage")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return nil, err
	}

	if inputFlags.Help {
		flags.PrintDefaults()
		os.Exit(1)
	}

	return &inputFlags, nil
}
