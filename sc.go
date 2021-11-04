package main

import (
	"fmt"
	clientmodel "github.com/filswan/go-swan-client/model"
	"github.com/filswan/go-swan-client/subcommand"
)

// type ConfCar struct {
// 	LotusApiUrl      string
// 	LotusAccessToken string
// 	OutputDir        string
// 	InputDir         string
// }

func main() {
	confCar := &clientmodel.ConfCar{
		LotusApiUrl:        "http://192.168.88.41:1234/rpc/v0",
		LotusAccessToken:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.-Y4pF34RGOten6YXoau-sEMOWOEeiHwGh9u2lsl4cv8",
		OutputDir:          "/home/gh/go/src/dev/test/deal/car",
		InputDir:           "/home/gh/go/src/dev/test/deal/source",
		GocarFileSizeLimit: 8589934592,
	}

	fmt.Println(confCar)

	fileDesc,_ :=subcommand.CreateCarFiles(confCar)
	fmt.Println(fileDesc)
}
