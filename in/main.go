package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tscolari/concourse-datadog-event-resource/params"
	datadog "github.com/zorkian/go-datadog-api"
)

func main() {
	input, err := params.ParseInput(os.Stdin)
	must(err)

	storePath := os.Args[1]

	client := datadog.NewClient(input.Source.ApiKey, input.Source.ApplicationKey)
	event, err := client.GetEvent(input.Version.Id)
	must(err)

	eventFile, err := os.Create(filepath.Join(storePath, "event.json"))
	must(err)

	must(json.NewEncoder(eventFile).Encode(event))

	response := params.InOutResponse{
		Version:  params.Version{Id: event.Id},
		Metadata: params.EventToMetadata(*event),
	}

	must(json.NewEncoder(os.Stdout).Encode(response))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
