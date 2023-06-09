package main

import (
	"github.com/provokateurin/rain-cloud/pkg/common"
	_ "github.com/provokateurin/rain-cloud/pkg/include"
)

func main() {
	err := common.Serve(8080)
	if err != nil {
		panic(err)
	}
}
