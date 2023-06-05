package main

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
)

func main() {
	err := common.Serve(8000)
	if err != nil {
		panic(err)
	}
}
