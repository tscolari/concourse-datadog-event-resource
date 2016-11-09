package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/tscolari/concourse-datadog-event-resource/resource"
	datadog "github.com/zorkian/go-datadog-api"
)

func main() {
	input, err := resource.ParseInput(os.Stdin)

	storePath := os.Args[1]

	client := datadog.NewClient(input.Source.ApiKey, input.Source.ApplicationKey)
	ddResource := resource.NewDatadog(client)

	response, event, err := ddResource.In(input)
	must(err)

	eventFile, err := os.Create(filepath.Join(storePath, "event.json"))
	must(err)

	must(json.NewEncoder(eventFile).Encode(event))
	must(json.NewEncoder(os.Stdout).Encode(response))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
