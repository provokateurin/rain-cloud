package main

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func main() {
	err := common.Serve(8080)
	if err != nil {
		panic(err)
	}
}
