# cowin #
This programs pings Cowin API's at a given interval to check if any vaccine slots are available. We can search via districts and provide multiple filters.

## Usage ##
`go run cmd/main.go --help`
```
Usage of cowin:
  -districts string
        <district>,<district>,... eg: 143,150
  -pincodes string
        <pincode>,<pincode>,... eg: 110046,110047
  -filter value
        <dose type>,<minimum age>,<vaccine> eg: 1,18,1 [Dose Type: 1st dose (1), 2nd dose (2)] [Vaccine: COVAXIN (1), COVISHIELD (2)] [Required]
  -interval string
        Interval [1s/2m/3h] (default "5m")
  -min-alert-value int
        Minimum doses required for sending alert message. (default 5)
  -help
        Set to true for printing usage
```

## Run ##
`go run cmd/main.go --pincodes 110046,110088 --filter 1,18,1 --filter 1,18,2` 
## Build and Run ##
#### linux / macOS ####
* `go build -o cowin cmd/main.go`
* `./cowin --pincodes 110046,110088 --filter 1,18,1 --filter 1,18,2`
#### Windows ####
* `go build -o cowin.exe cmd/main.go`
* `cowin.exe --pincodes 110046,110088 --filter 1,18,1 --filter 1,18,2` 

##### Examples #####
* `go run cmd/main.go --pincodes 110046,110088 --filter 1,18,1 --filter 1,18,2`
* `go run cmd/main.go --districts 150,143 --filter 1,18,1 --filter 1,18,2`
* `go run cmd/main.go --pincodes 110046,110088 --districts 150,143 --filter 1,18,2`
* `go run cmd/main.go --districts 150,143 --filter 1,18,1 --interval 10m --min-alert-value 5`


## Install Latest Go ##
https://golang.org/doc/install