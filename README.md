**cowin**
This programs pings Cowin API's at a given interval to check if any vaccine slots are available. We can search via districts and provide multiple filters.

*Usage*
`go run cmd/main.go --help`
```
  -districts string
        <district>,<district>,... eg: 143,150 
  -filter value
        <dose type>,<minimum age>,<vaccine> eg: 1,18,1 [Dose Type: 1st dose (1), 2nd dose (2)] [Vaccine: COVAXIN (1), COVISHIELD (2)]
  -help
        Set to true for printing usage
  -interval string
        Interval [1s/2m/3h] (default "5m")
```

*Build*
`go build -o cowin cmd/main.go`

*Run*
`go run cmd/main.go --districts 150,143 --filter 1,18,1 --filter 1,18,2 --interval 10m`
or
`./cowin --districts 150,143 --filter 1,18,1 --filter 1,18,2 --interval 10m`

