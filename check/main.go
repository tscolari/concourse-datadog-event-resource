package main

import (
	"encoding/json"
	"os"

	"github.com/tscolari/concourse-datadog-event-resource/resource"
	datadog "github.com/zorkian/go-datadog-api"
)

func main() {
	input, err := resource.ParseInput(os.Stdin)
	must(err)

	client := datadog.NewClient(input.Source.ApiKey, input.Source.ApplicationKey)
	ddResource := resource.NewDatadog(client)

	response, err := ddResource.Check(input)
	must(err)
	must(json.NewEncoder(os.Stdout).Encode(response))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
