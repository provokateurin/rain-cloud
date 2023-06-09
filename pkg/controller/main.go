package main

import (
	v1 "github.com/provokateurin/rain-cloud/pkg/controller/api/v1"
)

func main() {
	err := v1.Serve(8080)
	if err != nil {
		panic(err)
	}
}
