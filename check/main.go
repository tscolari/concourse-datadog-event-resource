package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/tscolari/concourse-datadog-event-resource/params"
	datadog "github.com/zorkian/go-datadog-api"
)

func main() {
	input, err := params.ParseInput(os.Stdin)
	must(err)

	client := datadog.NewClient(input.Source.ApiKey, input.Source.ApplicationKey)
	end := time.Now()
	start := end.AddDate(0, -2, 0)

	events, err := client.GetEvents(
		int(start.Unix()),
		int(end.Unix()),
		input.Source.Priority,
		input.Sources(),
		input.Tags(),
	)
	must(err)

	checkResponse := params.CheckResponse{}
	for _, event := range events {
		checkResponse = append(checkResponse, params.Version{Id: event.Id})
	}

	must(json.NewEncoder(os.Stdout).Encode(checkResponse))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
